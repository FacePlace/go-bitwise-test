package bitwise

import (
	"testing"
)

func TestGetInt(t *testing.T) {
	sample := []byte{84, 69, 83, 52, 54, 00, 00, 00}

	var test1 int8 = 84
	var test2 int16 = 17748
	var test3 int32 = 877872468
	var test4 int64 = 232806106452

	var test5 uint8 = 84
	var test6 uint16 = 17748
	var test7 uint32 = 877872468
	var test8 uint64 = 232806106452

	actual1 := GetInt8(sample)
	if test1 != actual1 {
		t.Errorf("Expected int8 to be %v, got - %v", test1, actual1)
	}

	actual2 := GetInt16(sample)
	if test2 != actual2 {
		t.Errorf("Expected int16 to be %v, got - %v", test2, actual2)
	}

	actual3 := GetInt32(sample)
	if test3 != actual3 {
		t.Errorf("Expected int32 to be %v, got - %v", test3, actual3)
	}

	actual4 := GetInt64(sample)
	if test4 != actual4 {
		t.Errorf("Expected int64 to be %v, got - %v", test4, actual4)
	}

	actual5 := GetUint8(sample)
	if test5 != actual5 {
		t.Errorf("Expected uint8 to be %v, got - %v", test5, actual5)
	}

	actual6 := GetUint16(sample)
	if test6 != actual6 {
		t.Errorf("Expected uint16 to be %v, got - %v", test6, actual6)
	}

	actual7 := GetUint32(sample)
	if test7 != actual7 {
		t.Errorf("Expected uint32 to be %v, got - %v", test7, actual7)
	}

	actual8 := GetUint64(sample)
	if test8 != actual8 {
		t.Errorf("Expected uint64 to be %v, got - %v", test8, actual8)
	}
}
