package main

import (
	"context"
	"encoding/json"
	"errors"
	"fib_grpc/proto"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var rds *redis.Client

func init() {
	fibInit := initFibArray()

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPort := os.Getenv("REDIS_PORT")
	rds = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisAddr, redisPort),
		Password: "",
		DB:       0,
	})

	err := writeData(fibInit)
	if err != nil {
		log.Println(err)
	}
}

func initFibArray() []uint64 {
	initArray := make([]uint64, 50)
	initArray[0], initArray[1] = 0, 1
	for j := 2; j < len(initArray); j++ {
		initArray[j] = initArray[j-1] + initArray[j-2]
	}

	return initArray
}

func writeData(fibData []uint64) error {
	ctx := context.Background()

	jsonFibData, err := json.Marshal(fibData)
	if err != nil {
		return err
	}

	status := rds.Set(ctx, "fib", string(jsonFibData), 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func readData() ([]uint64, error) {
	val, err := rds.Get(context.Background(), "fib").Result()
	switch {
	case err == redis.Nil:
		log.Println("key does not exist")
		return nil, err
	case err != nil:
		log.Println("get failed: ", err)
		return nil, err
	case val == "":
		err = errors.New("value is empty")
		log.Println(err)
		return nil, err
	}
	var fibSlice []uint64
	json.Unmarshal([]byte(val), &fibSlice)

	return fibSlice, nil
}

type fibonacciServer struct {
	proto.UnimplementedFibonacciServer
}

func (s *fibonacciServer) GetFibonacci(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("[Fibonacci grpc]")
	fibSlice, err := getFibSliceByIndexes(request.GetStart(), request.GetEnd())
	if err != nil {
		log.Println("error in getting slice: ", err)
		return nil, err
	}

	jsonArray, err := json.Marshal(fibSlice)

	return &proto.Response{Fibslice: string(jsonArray)}, nil
}

func HttpFibAction(w http.ResponseWriter, r *http.Request) {
	log.Println("[Fib action]")
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

	fibSlice, err := getFibSliceByIndexes(int32(start), int32(end))
	if err != nil {
		log.Println("error in getting slice: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(convertIntArrayToStr(fibSlice)))
}

func getFibSliceByIndexes(start, end int32) ([]uint64, error) {
	fibArr, err := readData()
	if err != nil {
		return nil, err
	}

	if start < 1 {
		start = 1
	}

	if start > end {
		end = start
	}

	if int(end) > len(fibArr) {
		fibArr = calculateToNewEnd(fibArr, int(end))
		go writeData(fibArr)
	}

	return fibArr[start-1:end], nil
}

func calculateToNewEnd(fibArray []uint64, newEnd int) []uint64 {
	newFibArray := make([]uint64, 0, newEnd)
	newFibArray = append(newFibArray, fibArray...)
	for i := len(fibArray); i < newEnd; i++ {
		newFibArray = append(newFibArray, newFibArray[i-1] + newFibArray[i-2])
	}
	return newFibArray
}

func convertIntArrayToStr(values []uint64) string {
	valuesText := make([]string, 0, len(values))
	for i := range values {
		number := values[i]
		text := strconv.FormatUint(number, 10)
		valuesText = append(valuesText, text)
	}

	return fmt.Sprintf("[%s]", strings.Join(valuesText, ", "))
}

func newServer() *fibonacciServer {
	s := &fibonacciServer{}
	return s
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

func startHttpListener(errChan chan<-error) {
	http.HandleFunc("/get", HttpFibAction)
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