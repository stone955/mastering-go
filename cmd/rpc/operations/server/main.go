package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
	"time"

	"mastering-go/cmd/rpc/operations/protocol"
)

// Operations 服务
type Operations struct{}

// Multiply 乘法
func (opt *Operations) Multiply(args *protocol.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Multiply 乘法
func (opt *Operations) Divide(args *protocol.Args, quotient *protocol.Quotient) error {
	if args.B == 0 {
		return errors.New("args.B must not be 0")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.A % args.B
	return nil
}

func main() {
	srv := rpc.NewServer()

	// 注册服务
	if err := srv.Register(&Operations{}); err != nil {
		log.Fatal(err)
	}

	// 通过 http 监听，到时做协议转换
	if err := http.ListenAndServe(":3000", srv); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	log.Printf("rpc server close at %v\n", time.Now())
}
