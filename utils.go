package main

func getInt8(buffer []byte) (n int8) {
	nBytes := 1

	_ = buffer[nBytes-1]

	i := 0
	for i < nBytes {
		n |= int8(buffer[i]) << (8 * (i))

		i++
	}

	return
}

func getInt16(buffer []byte) (n int16) {
	nBytes := 2

	_ = buffer[nBytes-1]

	i := 0
	for i < nBytes {
		n |= int16(buffer[i]) << (8 * (i))

		i++
	}

	return
}

func getInt32(buffer []byte) (n int32) {
	nBytes := 4

	_ = buffer[nBytes-1]

	i := 0
	for i < nBytes {
		n |= int32(buffer[i]) << (8 * (i))

		i++
	}

	return
}

func getInt64(buffer []byte) (n int64) {
	nBytes := 8

	_ = buffer[nBytes-1]

	i := 0
	for i < nBytes {
		n |= int64(buffer[i]) << (8 * (i))

		i++
	}

	return
}
