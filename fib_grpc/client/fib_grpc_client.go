package main

import (
	"context"
	"errors"
	"fib_grpc/proto"
	"google.golang.org/grpc"
	"log"
	"math/big"
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
		time.Sleep(500 * time.Microsecond)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			start, end := rand.Int31n(150), rand.Int31n(160)
			fibSlice, err := client.GetFibonacci(ctx, &proto.Request{Start: start, End: end})
			if err != nil {
				log.Println(err)
				return
			}
			slice, err := castSliceToBigInt(fibSlice.GetSlice())
			if err != nil {
				log.Println(err)
			}
			log.Println(slice)
		}(wg)
	}
	wg.Wait()
}

func castSliceToBigInt(slice []*proto.BigInt) ([]*big.Int, error) {
	fibSlice := make([]*big.Int, len(slice))
	for i := range slice {
		elem, ok := new(big.Int).SetString(string(slice[i].GetBigInt()), 0)
		if !ok {
			return nil, errors.New("error in casting bigInts")
		}
		fibSlice[i] = elem
	}
	return fibSlice, nil
}
