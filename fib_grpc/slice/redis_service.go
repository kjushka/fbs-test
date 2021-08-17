package slice

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
	"strconv"
)

func writeData(fibSlice []*big.Int, key int32) error {
	ctx := context.Background()

	jsonFibData, err := json.Marshal(fibSlice)
	if err != nil {
		return err
	}

	keyStr := strconv.FormatInt(int64(key), 10)

	status := rds.Set(ctx, keyStr, string(jsonFibData), 0)
	if status.Err() != nil {
		return status.Err()
	}

	log.Println("data written with length: ", len(fibSlice))
	return nil
}

func readData(key int32) ([]*big.Int, error) {
	keyStr := strconv.FormatInt(int64(key), 10)
	val, err := rds.Get(context.Background(), keyStr).Result()
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
	var fibSlice []*big.Int
	err = json.Unmarshal([]byte(val), &fibSlice)
	if err != nil {
		return nil, err
	}

	log.Println("data read with length: ", len(fibSlice))
	return fibSlice, nil
}
