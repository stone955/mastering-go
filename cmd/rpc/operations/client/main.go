package main

import (
	"log"
	"net/rpc"

	"mastering-go/cmd/rpc/operations/protocol"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	dones := make([]chan *rpc.Call, 0, 10)

	log.Println("########## 同步发起请求开始 ##########")

	// 同步发起请求
	for i := 0; i < 50; i++ {
		quotient := &protocol.Quotient{}
		args := &protocol.Args{
			A: i * 2,
			B: i + 5,
		}

		divCall := client.Go("Operations.Divide", args, quotient, nil)
		dones = append(dones, divCall.Done)
	}

	log.Println("########## 同步发起请求结束 ##########")

	log.Println("########## 异步读取响应开始 ##########")
	// 异步读取响应
	for idx, done := range dones {
		replyCall := <-done
		args := replyCall.Args.(*protocol.Args)
		reply := replyCall.Reply.(*protocol.Quotient)
		log.Printf("receive %v: %v / %v = %v, %v %% %v = %v\n",
			idx, args.A, args.B, reply.Quo, args.A, args.B, reply.Rem)
	}

	log.Println("########## 异步读取响应结束 ##########")
}
