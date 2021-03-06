package main

import (
	"context"
	"fib_grpc/proto"
	"fib_grpc/slice"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
)

type fibonacciServer struct {
	proto.UnimplementedFibonacciServer
}

func (s *fibonacciServer) GetFibonacci(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("[Fibonacci grpc]")
	fibSlice, err := slice.GetFibSliceByIndexes(ctx, request.GetStart(), request.GetEnd())
	if err != nil {
		log.Println("error in getting slice: ", err)
		return nil, err
	}

	grpcSlice, err := slice.CastSliceToGrpcType(fibSlice)
	if err != nil {
		return nil, err
	}

	return &proto.Response{Slice: grpcSlice}, nil
}

func startGrpcListener(errChan chan<- error) {
	listener, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterFibonacciServer(grpcServer, newServer())
	log.Println("Grpc listen localhost:8090")
	err = grpcServer.Serve(listener)
	errChan <- err
}

func newServer() *fibonacciServer {
	s := &fibonacciServer{}
	return s
}

func GetFibonacciHttp(w http.ResponseWriter, r *http.Request) {
	log.Println("[Fib http]")
	query := r.URL.Query()
	start, err := strconv.Atoi(query.Get("start"))
	if err != nil {
		log.Println("error in atoi start")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	end, err := strconv.Atoi(query.Get("end"))
	if err != nil {
		log.Println("error in atoi end")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fibSlice, err := slice.GetFibSliceByIndexes(ctx, int32(start), int32(end))
	if err != nil {
		log.Println("error in getting slice: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(slice.ConvertIntArrayToStr(fibSlice)))
	if err != nil {
		http.Error(w, "Error in writing slice", http.StatusInternalServerError)
	}
}

func startHttpListener(errChan chan<- error) {
	http.HandleFunc("/fib", GetFibonacciHttp)
	log.Println("Http listen localhost:8070")
	err := http.ListenAndServe(":8070", nil)
	errChan <- err
}

func main() {
	errChan := make(chan error)
	go startGrpcListener(errChan)
	go startHttpListener(errChan)

	defer close(errChan)
	for range errChan {
		log.Fatal(errChan)
	}
}
