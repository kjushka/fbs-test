package slice

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)

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
	err = json.Unmarshal([]byte(val), &fibSlice)
	if err != nil {
		return nil, err
	}
	return fibSlice, nil
}