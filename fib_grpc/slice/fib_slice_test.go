package slice

import (
	"math/big"
	"testing"
)

func TestConvertIntArrayToStr(t *testing.T) {
	initSlice := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(4),
		big.NewInt(5),
		big.NewInt(6),
		big.NewInt(7),
		big.NewInt(8),
		big.NewInt(9),
		big.NewInt(10),
	}
	strInitSlice := ConvertIntArrayToStr(initSlice)
	if strInitSlice != "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]" {
		t.Errorf("Converter with filled array test failed")
	}

	var emptySlice []*big.Int
	strEmptySlice := ConvertIntArrayToStr(emptySlice)
	if strEmptySlice != "[]" {
		t.Errorf("Converter with empty array test failed")
	}
}

func bigIntSliceEquals(a, b []*big.Int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Cmp(b[i]) != 0 {
			return false
		}
	}
	return true
}

func TestInitFibArray(t *testing.T) {
	initFibArrayWith2 := InitFibArray(2)
	t.Log(initFibArrayWith2)
	compare := []*big.Int{big.NewInt(0), big.NewInt(1)}
	if !bigIntSliceEquals(initFibArrayWith2, compare) {
		t.Errorf("Init with size 2 failed")
	}

	initFibArrayWith0 := InitFibArray(0)
	t.Log(initFibArrayWith0)
	if !bigIntSliceEquals(initFibArrayWith0, compare) {
		t.Errorf("Init with size 2 failed")
	}

	initFibArrayWith5 := InitFibArray(5)
	t.Log(initFibArrayWith5)
	compare = []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
	}
	if !bigIntSliceEquals(initFibArrayWith5, compare) {
		t.Errorf("Init with size 5 failed")
	}
}

func TestCalculateToNewEnd(t *testing.T) {
	initFibArray := InitFibArray(5)
	newFibArrayWithEnd8 := CalculateToNewEnd(initFibArray, 8)
	compare := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(5),
		big.NewInt(8),
		big.NewInt(13),
	}
	if !bigIntSliceEquals(newFibArrayWithEnd8, compare) {
		t.Errorf("Calculate new end array with end 8 failed")
	}

	compare = []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
	}
	newFibArrayWithEnd4 := CalculateToNewEnd(initFibArray, 4)
	if !bigIntSliceEquals(newFibArrayWithEnd4, compare) {
		t.Errorf("Calculate new end array with end 4 failed")
	}
}
