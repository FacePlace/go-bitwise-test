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

func getUint8(buffer []byte) (n uint8) {
	nBytes := 1
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint8(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func getUint16(buffer []byte) (n uint16) {
	nBytes := 2
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint16(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func getUint32(buffer []byte) (n uint32) {
	nBytes := 4
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint32(buffer[i]) << (8 * (i))
		i++
	}
	return
}

func getUint64(buffer []byte) (n uint64) {
	nBytes := 8
	_ = buffer[nBytes-1]
	i := 0
	for i < nBytes {
		n |= uint64(buffer[i]) << (8 * (i))
		i++
	}
	return
}
