package slice

/*func TestConvertIntArrayToStr(t *testing.T) {
	initSlice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	strInitSlice := ConvertIntArrayToStr(initSlice)
	if strInitSlice != "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]" {
		t.Errorf("Converter with filled array test failed")
	}

	var emptySlice []uint64
	strEmptySlice := ConvertIntArrayToStr(emptySlice)
	if strEmptySlice != "[]" {
		t.Errorf("Converter with empty array test failed")
	}
}

func uint64SliceEquals(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInitFibArray(t *testing.T) {
	initFibArrayWith2 := InitFibArray(2)
	t.Log(initFibArrayWith2)
	if compare := []uint64{0, 1}; !uint64SliceEquals(initFibArrayWith2, compare) {
		t.Errorf("Init with size 2 failed")
	}

	initFibArrayWith0 := InitFibArray(0)
	t.Log(initFibArrayWith0)
	if compare := []uint64{0, 1}; !uint64SliceEquals(initFibArrayWith0, compare) {
		t.Errorf("Init with size 2 failed")
	}

	initFibArrayWith5 := InitFibArray(5)
	t.Log(initFibArrayWith5)
	if compare := []uint64{0, 1, 1, 2, 3}; !uint64SliceEquals(initFibArrayWith5, compare) {
		t.Errorf("Init with size 2 failed")
	}
}

func TestCalculateToNewEnd(t *testing.T) {
	initFibArray := InitFibArray(5)
	newFibArrayWithEnd8 := CalculateToNewEnd(initFibArray, 8)
	if compare := []uint64{0, 1, 1, 2, 3, 5, 8, 13}; !uint64SliceEquals(newFibArrayWithEnd8, compare) {
		t.Errorf("Calculate new end array with end 8 failed")
	}

	newFibArrayWithEnd4 := CalculateToNewEnd(initFibArray, 4)
	if compare := []uint64{0, 1, 1, 2, 3}; !uint64SliceEquals(newFibArrayWithEnd4, compare) {
		t.Errorf("Calculate new end array with end 4 failed")
	}
}*/
