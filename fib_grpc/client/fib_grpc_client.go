package main

import (
	"context"
	"fib_grpc/proto"
	"fib_grpc/slice"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"sync"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		time.Sleep(2000 * time.Microsecond)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			start, end := rand.Int31n(150), rand.Int31n(160)
			fibSlice, err := client.GetFibonacci(ctx, &proto.Request{Start: start, End: end})
			if err != nil {
				log.Println(err)
				return
			}
			grpcSlice, err := slice.CastSliceToBigInt(fibSlice.GetSlice())
			if err != nil {
				log.Println(err)
			}
			log.Println(grpcSlice)
		}(wg)
	}
	wg.Wait()
}
