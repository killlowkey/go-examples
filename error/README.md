# 错误处理最佳实践

使用 github.com/pkg/errors 包来处理错误，这个包提供了一些有用的函数来处理错误，比如 `errors.Wrap` 和 `errors.Cause`
。这个包的使用方法是，当一个函数返回一个错误时，调用 `errors.Wrap`
函数，这个函数会将当前的错误包装成一个新的错误，同时记录当前的调用栈信息。当调用 `errors.Cause`
函数时，会返回原始的错误。这样做的好处是，可以在程序的最顶层打印出完整的调用栈信息，方便定位问题。

1. errors.New: 创建新的错误，包含调用栈信息, 用于内部错误
2. errors.Wrap: 包装第三方库错误，添加调用栈信息
3. errors.Cause: 获取原始错误

## Errors 原生库
> 使用 github.com/pkg/errors 来代替

1. 原生库不携带 stacktrace 
2. 原生库不支持 Wrap

## Reference

1. [Go Error 的处理最佳实践](https://chanjarster.github.io/post/go/err-throw-rules/)
   1. 自己 new 的 error，根据情况包含 stacktrace
   2. 不要 wrap 自己代码返回的 error 
   3. wrap 第三方库返回的 error 
   4. 尽量只把 error 用作异常情况
2. [Go错误处理最佳实践](https://lailin.xyz/post/go-training-03.html)
3. [(我的) Golang 错误处理最佳实践](https://xuanwo.io/2020/05-go-error-handling/)
4. [Go2 草稿提案](https://go.googlesource.com/proposal/+/master/design/go2draft.md)
   > Go 2 error handling proposal, 将 github.com/pkg/errors 的功能集成到标准库中