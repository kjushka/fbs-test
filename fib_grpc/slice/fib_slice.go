package slice

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"strings"
)

var rds *redis.Client

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

	err = writeData(fibInit)
	if err != nil {
		log.Println(err)
	}
}

func InitFibArray(end int) []uint64 {
	if end < 2 {
		end = 2
	}
	initArray := make([]uint64, end)
	initArray[0], initArray[1] = 0, 1
	for j := 2; j < len(initArray); j++ {
		initArray[j] = initArray[j-1] + initArray[j-2]
	}

	return initArray
}

func GetFibSliceByIndexes(start, end int32) ([]uint64, error) {
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
		fibArr = CalculateToNewEnd(fibArr, int(end))
		go writeData(fibArr)
	}

	return fibArr[start-1 : end], nil
}

func CalculateToNewEnd(fibArray []uint64, newEnd int) []uint64 {
	if newEnd <= len(fibArray) {
		newArray := make([]uint64, 0, len(fibArray))
		return append(newArray, fibArray...)
	}
	newFibArray := make([]uint64, 0, newEnd)
	newFibArray = append(newFibArray, fibArray...)
	for i := len(fibArray); i < newEnd; i++ {
		newFibArray = append(newFibArray, newFibArray[i-1]+newFibArray[i-2])
	}
	return newFibArray
}

func ConvertIntArrayToStr(values []uint64) string {
	if len(values) == 0 {
		return "[]"
	}

	valuesText := make([]string, 0, len(values))
	for i := range values {
		number := values[i]
		text := strconv.FormatUint(number, 10)
		valuesText = append(valuesText, text)
	}

	return fmt.Sprintf("[%s]", strings.Join(valuesText, ", "))
}
