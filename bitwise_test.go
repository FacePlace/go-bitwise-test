package bitwise

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, test1, actual1)

	actual2 := GetInt16(sample)
	assert.Equal(t, test2, actual2)

	actual3 := GetInt32(sample)
	assert.Equal(t, test3, actual3)

	actual4 := GetInt64(sample)
	assert.Equal(t, test4, actual4)

	actual5 := GetUint8(sample)
	assert.Equal(t, test5, actual5)

	actual6 := GetUint16(sample)
	assert.Equal(t, test6, actual6)

	actual7 := GetUint32(sample)
	assert.Equal(t, test7, actual7)

	actual8 := GetUint64(sample)
	assert.Equal(t, test8, actual8)
}
