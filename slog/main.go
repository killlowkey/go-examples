package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type ctxKey string

const (
	slogFields ctxKey = "slog_fields"
)

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds a slog attribute to the provided context so that it will be
// included in any Record created with such context
func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogFields, v)
	}

	var v []slog.Attr
	v = append(v, attr)
	return context.WithValue(parent, slogFields, v)
}

// main https://betterstack.com/community/guides/logging/logging-in-go/
// https://github.com/samber/slog-gin
// https://betterstack.com/logs
// slog 提案：https://go.googlesource.com/proposal/+/master/design/56345-structured-logging.md
func main() {
	flog, err := os.OpenFile("F:\\go-examples\\slog\\logs\\app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("error opening file: %v", err))
	}

	// 日志级别可以控制日志的输出
	// 输出每个级别可输出的日志类型, 级别值越小, 输出的日志类型越多，例如：
	// slog.LevelDebug 输出 Debug, Info, Warn, Error 级别的日志
	// slog.LevelInfo 输出 Info, Warn, Error 级别的日志
	// slog.LevelWarn 输出 Warn, Error 级别的日志
	// slog.LevelError 只输出 Error 级别的日志
	options := &slog.HandlerOptions{Level: slog.LevelInfo}
	// logger := slog.New(slog.NewTextHandler(io.MultiWriter(os.Stdout, flog), options))
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, flog), options))
	slog.SetDefault(logger)

	slog.Info("I am default slog logger")

	logger.With("requestId", "f93eb275-fc39-4556-ad7d-94bbecb0c7bc").Info("Hello, world!")
	logger.Error("This is an error message", slog.String("error", "something went wrong"))
	logger.InfoContext(context.TODO(), "This is a message with context")

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"This is a message with attributes",
		slog.String("key", "value"),
	)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)

	requestLog := slog.Default().With("requestId", "f93eb275-fc39-4556-ad7d-94bbecb0c7bc")
	requestLog.Info("read file form client")

	// Using the context package with Slog
	// https://betterstack.com/community/guides/logging/logging-in-go/#using-the-context-package-with-slog
	// 需要扩展 slog.Handler 接口，以便在记录日志之前将属性添加到记录中
	h := &ContextHandler{slog.NewJSONHandler(os.Stdout, nil)}
	logger = slog.New(h)
	ctx := AppendCtx(context.Background(), slog.String("request_id", "req-123"))
	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))

}
