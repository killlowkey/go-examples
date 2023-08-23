package common

import "context"

// Args 方法参数
type Args struct {
	A int
	B int
}

// Reply 响应
type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
