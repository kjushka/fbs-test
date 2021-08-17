package slice

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

var rds *redis.Client
var maxLen int32

func init() {
	initSize, err := strconv.Atoi(os.Getenv("INIT_SIZE"))
	if err != nil || initSize < 2 {
		log.Println("Init size should be number and more than 1")
		initSize = 2
	}

	fibInit := InitFibArray(initSize)

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPort := os.Getenv("REDIS_PORT")
	rds = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisAddr, redisPort),
		Password: "",
		DB:       0,
	})

	atomic.StoreInt32(&maxLen, int32(initSize))
	err = writeData(fibInit, int32(initSize))
	if err != nil {
		log.Println(err)
	}
}

func InitFibArray(end int) []*big.Int {
	if end < 2 {
		end = 2
	}
	initArray := make([]*big.Int, end)
	initArray[0], initArray[1] = big.NewInt(0), big.NewInt(1)
	for j := 2; j < len(initArray); j++ {
		newElem := new(big.Int)
		newElem.Add(initArray[j-1], initArray[j-2])
		initArray[j] = newElem
	}

	return initArray
}

func GetFibSliceByIndexes(start, end int32) ([]*big.Int, error) {
	fibArr, err := readData(atomic.LoadInt32(&maxLen))
	if err != nil {
		return nil, err
	}

	if start < 0 {
		start = 0
	}

	if start > end {
		end = start
	}

	if int(end) >= len(fibArr) {
		fibArr = CalculateToNewEnd(fibArr, int(end+1))
		if end+1 > atomic.LoadInt32(&maxLen) {
			err = writeData(fibArr, maxLen)
			if err != nil {
				return nil, err
			} else {
				atomic.StoreInt32(&maxLen, end+1)
			}
		}
	}

	return fibArr[start : end+1], nil
}

func CalculateToNewEnd(fibArray []*big.Int, newEnd int) []*big.Int {
	if newEnd <= len(fibArray) {
		newArray := make([]*big.Int, 0, len(fibArray))
		return append(newArray, fibArray...)
	}
	newFibArray := make([]*big.Int, 0, newEnd)
	newFibArray = append(newFibArray, fibArray...)
	for i := len(fibArray); i < newEnd; i++ {
		newFibArray = append(newFibArray, new(big.Int).Add(newFibArray[i-1], newFibArray[i-2]))
	}
	return newFibArray
}

func ConvertIntArrayToStr(values []*big.Int) string {
	if len(values) == 0 {
		return "[]"
	}

	valuesText := make([]string, 0, len(values))
	for _, e := range values {
		text := e.Text(10)
		valuesText = append(valuesText, text)
	}

	return fmt.Sprintf("[%s]", strings.Join(valuesText, ", "))
}
