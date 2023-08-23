package main

import "zap-example/log"

func main() {
	o := &log.Options{
		Level:             "debug",
		Format:            "console",
		OutputPaths:       []string{"tmp/app.log", "stdout"},
		DisableStacktrace: false,
		DisableCaller:     false,
	}

	log.Init(o)
	log.Infow("hello world", "name", "zap")
	log.Debugw("bind address: ", "host", "127.0.0.1", "port", 8080)
	log.Infof("args1=%d arg2=%d", 1, 2)
}
