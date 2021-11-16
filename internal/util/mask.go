package util

import "fmt"

// MaskBits64 generates a bitmask with n of the lowest bits set.
func MaskBits64(n uint) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n > 64 {
		panic(fmt.Sprintf("%d bits exceeds maximum permitted value of 64", n))
	}

	value := uint64(1)
	for i := uint(1); i < n; i++ {
		value = value | (value << 1)
	}
	return value
}
