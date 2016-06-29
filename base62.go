package base62

import (
	"bytes"
	"math"
)

// BASE62 should always be 0-9,a-z,A-Z
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = len(alphabet)

var balphabet = []byte(alphabet)

// Encode a positive integer to a Base62 token
func Encode(num int) []byte {
	if num == 0 {
		return []byte{alphabet[0]}
	}
	// result byte arr
	var res []byte
	// remainder of modulo
	var rem int
	// calculate base (should be 62)

	// until num is == 0
	for num != 0 {
		// calculate remainder
		rem = num % base
		// calculate quotient
		num = num / base
		// push alphabet[rem] into res
		res = append(res, alphabet[rem])
	}
	// reverse res
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// Decode a Base62 token to its original positive integer
func Decode(str []byte) int {
	// calculate base (should be 62)
	strlen := len(str)
	num := 0
	idx := 0
	// until num is == 0
	for _, c := range str {
		// calculate remainder
		power := (strlen - (idx + 1))
		// calculate quotient
		index := bytes.IndexByte(balphabet, c)
		num += index * int(math.Pow(float64(base), float64(power)))
		idx++
	}

	return num
}
