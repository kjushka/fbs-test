package main

import (
	"context"
	"fib_grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:8090", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := proto.NewFibonacciClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	fibSlice, err := client.GetFibonacci(ctx, &proto.Request{Start: 0, End: 62})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(fibSlice.Fibslice)
}
