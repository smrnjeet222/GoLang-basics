package main

import (
	"fmt"
)

const (
	_  = iota             // ignoring th 1st value
	kb = 1 << (10 * iota) // 2^10
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

func convFileSize() {
	fileSize := 540000000000.
	fmt.Printf("%.3f GB", fileSize/gb)
}
