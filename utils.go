package main

import "fmt"

func getInt32(buffer []byte) (n int32) {
	_ = buffer[3]
	n |= int32(buffer[0])
	fmt.Printf("%b\n", n)
	n |= int32(buffer[1]) << 8
	fmt.Printf("%b\n", n)
	n |= int32(buffer[2]) << 16
	fmt.Printf("%b\n", n)
	n |= int32(buffer[3]) << 24
	fmt.Printf("%b\n", n)
	return
}
