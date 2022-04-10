package main

import "fmt"

func main() {

	var x int8
	var y int16
	var z int32
	var w int64

	var x1 uint8
	var y1 uint16
	var z1 uint32
	var w1 uint64

	var x2 float32
	var y2 float64

	x = getInt8([]byte{84})
	fmt.Println(x)
	y = getInt16([]byte{84, 69})
	fmt.Println(y)
	z = getInt32([]byte{84, 69, 83, 52})
	fmt.Println(z)
	w = getInt64([]byte{84, 69, 83, 52, 54, 0, 0, 0})
	fmt.Println(w)

	x1 = getUint8([]byte{84})
	fmt.Println(x1)
	y1 = getUint16([]byte{84, 69})
	fmt.Println(y1)
	z1 = getUint32([]byte{84, 69, 83, 52})
	fmt.Println(z1)
	w1 = getUint64([]byte{84, 69, 83, 52, 54, 0, 0, 0})
	fmt.Println(w1)

	x2 = getFloat32([]byte{84, 69, 83, 52})
	fmt.Println(x2)
	y2 = getFloat64([]byte{84, 69, 83, 52, 54, 0, 0, 0})
	fmt.Println(y2)

}
