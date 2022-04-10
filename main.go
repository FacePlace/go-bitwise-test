package main

import "fmt"

func main() {

	var x int8
	var y int16
	var z int32
	var w int64

	x = getInt8([]byte{84})
	fmt.Println(x)
	y = getInt16([]byte{84, 69})
	fmt.Println(y)
	z = getInt32([]byte{84, 69, 83, 52})
	fmt.Println(z)
	w = getInt64([]byte{84, 69, 83, 52, 54, 0, 0, 0})
	fmt.Println(w)

}
