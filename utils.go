package main

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
