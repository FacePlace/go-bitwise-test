package bitwise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sample = []byte{84, 69, 83, 52, 54, 00, 00, 00}

func TestGetUint8(t *testing.T) {
	var expected uint8 = 84
	actual := GetUint8(sample)
	assert.Equal(t, expected, actual)
}

func TestGetUint16(t *testing.T) {
	var expected uint16 = 17748
	actual := GetUint16(sample)
	assert.Equal(t, expected, actual)
}

func TestGetUint32(t *testing.T) {
	var expected uint32 = 877872468
	actual := GetUint32(sample)
	assert.Equal(t, expected, actual)
}

func TestGetUint64(t *testing.T) {
	var expected uint64 = 232806106452
	actual := GetUint64(sample)
	assert.Equal(t, expected, actual)
}

func TestGetInt8(t *testing.T) {
	var expected int8 = 84
	actual := GetInt8(sample)
	assert.Equal(t, expected, actual)
}

func TestGetInt16(t *testing.T) {
	var expected int16 = 17748
	actual := GetInt16(sample)
	assert.Equal(t, expected, actual)
}

func TestGetInt32(t *testing.T) {
	var expected int32 = 877872468
	actual := GetInt32(sample)
	assert.Equal(t, expected, actual)
}

func TestGetInt64(t *testing.T) {
	var expected int64 = 232806106452
	actual := GetInt64(sample)
	assert.Equal(t, expected, actual)
}

func TestGetFloat32(t *testing.T) {
	var expected float32 = 1.9676127749335137e-7
	actual := GetFloat32(sample)
	assert.Equal(t, expected, actual)
}

func TestGetFloat64(t *testing.T) {
	var expected float64 = 1.1502149934e-312
	actual := GetFloat64(sample)
	assert.Equal(t, expected, actual)
}
