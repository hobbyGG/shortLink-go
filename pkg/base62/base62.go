package base62

import (
	"slices"
)

func Uint2string(seq uint64) string {
	var base string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	temp := make([]byte, 0, 8)
	for seq > 0 {
		i := seq % 62
		seq /= 62
		temp = append(temp, base[i])
	}
	slices.Reverse(temp)
	return string(temp)
}

var m = make(map[byte]int, 62)

func String2Uint(s string) uint64 {
	var base string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range 62 {
		m[base[i]] = i
	}
	bs := []byte(s)
	slices.Reverse(bs)
	var res uint64 = 0

	for n, char := range bs {
		res += uint64(m[char]) * pow(62, uint64(n))
	}
	return res
}

func pow(n, exp uint64) uint64 {
	if exp <= 0 {
		return 1
	}
	return n * pow(n, exp-1)
}
