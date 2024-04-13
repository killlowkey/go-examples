package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	// 使用 github.com/pkg/errors 构造的 error 则需要特殊的方式才能打印出 stacktrace：
	fmt.Printf("%+v", errors.New("abc"))

	// 通过 errors.Wrap 包装 error，可以在 error 中添加更多的信息
	if err := callThirtyLibrary(); err != nil {
		fmt.Printf("%+v", err)
	}

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	logger.Error("errors.New", zap.Error(errors.New("simple")))
	// 结果
	// {"level":"error","ts":"...","caller":"...","msg":"errors.New","err":"simple","errVerbose":"<stacktrace>"}

	logger.Error("errors.New", zap.NamedError("err", errors.New("simple")))
	// 结果和上面一样
}

// callThirtyLibrary 调用第三方库，出现错误，需要使用 errors.Wrap 包装错误
func callThirtyLibrary() error {
	// 模拟这是一个第三方库的调用
	f := func() error {
		return errors.New("third party library error")
	}

	return errors.Wrap(f(), "callThirtyLibrary")
}
