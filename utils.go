package bitwise

import "math"

func GetInt8(buffer []byte) (n int8) {
	nBytes := 1
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= int8(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetInt16(buffer []byte) (n int16) {
	nBytes := 2
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= int16(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetInt32(buffer []byte) (n int32) {
	nBytes := 4
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= int32(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetInt64(buffer []byte) (n int64) {
	nBytes := 8
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= int64(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetUint8(buffer []byte) (n uint8) {
	nBytes := 1
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint8(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetUint16(buffer []byte) (n uint16) {
	nBytes := 2
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint16(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetUint32(buffer []byte) (n uint32) {
	nBytes := 4
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint32(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetUint64(buffer []byte) (n uint64) {
	nBytes := 8
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint64(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func GetFloat32(buffer []byte) float32 {
	_ = buffer[3]
	n := GetUint32(buffer)

	return math.Float32frombits(n)
}

func GetFloat64(buffer []byte) float64 {
	_ = buffer[7]
	n := GetUint64(buffer)

	return math.Float64frombits(n)
}
