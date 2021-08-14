package main

/*func writeData(fibData []uint64) error {
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
}*/
