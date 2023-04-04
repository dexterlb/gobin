package gobin

// Appends n zero-bytes at the end of data and returns the new slice.
func PutBlank(n int, data []byte) []byte {
	return append(data, make([]byte, n)...)
}

// Appends the uint8 value to the end of data and returns the new slice
func PutUint8(v uint8, data []byte) []byte {
	return append(data, byte(v))
}

// Appends the little-endian uint16 value to the end of data and returns the new slice
func PutLUint16(v uint16, data []byte) []byte {
	return append(data, byte(v), byte(v>>8))
}

// Appends the big-endian uint16 value to the end of data and returns the new slice
func PutBUint16(v uint16, data []byte) []byte {
	return append(data, byte(v>>8), byte(v))
}

// AGetUint8 returns the uint8 value from the start of data and the remaining bytes
func AGetUint8(data []byte) (result uint8, tail []byte) {
	return GetUint8(data), data[1:]
}

// GetUint8 returns the uint8 value from the start of data
func GetUint8(data []byte) uint8 {
	return uint8(data[0])
}

// AGetLUint32 returns the little-endian uint32 value from the start of data and the remaining bytes
func AGetLUint32(data []byte) (result uint32, tail []byte) {
	return GetLUint32(data), data[4:]
}

// GetLUint32 returns the little-endian uint32 value from the start of data
func GetLUint32(data []byte) uint32 {
	return uint32(data[0]) + (uint32(data[1]) << 8) + (uint32(data[2]) << 16) + (uint32(data[3]) << 24)
}

// AGetBUint32 returns the big-endian uint32 value from the start of data and the remaining bytes
func AGetBUint32(data []byte) (result uint32, tail []byte) {
	return GetBUint32(data), data[4:]
}

// GetBUint32 returns the big-endian uint32 value from the start of data
func GetBUint32(data []byte) uint32 {
	return uint32(data[3]) + (uint32(data[2]) << 8) + (uint32(data[1]) << 16) + (uint32(data[0]) << 24)
}

// AGetLUint16 returns the little-endian uint16 value from the start of data and the remaining bytes
func AGetLUint16(data []byte) (result uint16, tail []byte) {
	return GetLUint16(data), data[2:]
}

// GetLUint16 returns the little-endian uint16 value from the start of data
func GetLUint16(data []byte) uint16 {
	return uint16(data[0]) + (uint16(data[1]) << 8)
}

// AGetBUint16 returns the big-endian uint16 value from the start of data and the remaining bytes
func AGetBUint16(data []byte) (result uint16, tail []byte) {
	return GetBUint16(data), data[2:]
}

// GetBUint16 returns the big-endian uint16 value from the start of data
func GetBUint16(data []byte) uint16 {
	return uint16(data[1]) + (uint16(data[0]) << 8)
}

// AGetString returns the UTF-8 encoded string with given length from the start
// of data and the remaining bytes
func AGetString(data []byte, length int) (result string, tail []byte) {
	return GetString(data, length), data[length:]
}

// GetString returns the UTF-8 encoded string with given length from the start of data
func GetString(data []byte, length int) string {
	i := 0
	for i < length {
		if data[i] == 0 {
			break
		}
		i++
	}

	return string(data[0:i])
}

// GetFlag returns the bit at the given bit offset in the first byte of given data
func GetFlag(data []byte, index uint8) bool {
	return ((data[0] >> index) & 0x1) == 0x1
}

// SetFlag sets the bit at the given bit offset in the first byte of given data of the first byte of given data
func SetFlag(data []byte, index uint8, value bool) {
	if value {
		data[0] |= 1 << index
	} else {
		data[0] &= ^(1 << index)
	}
}
