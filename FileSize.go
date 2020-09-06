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

// ConvFileSize for File conversion
func main() {
	fileSize := 540000000000.
	fmt.Printf("%.3f GB", fileSize/gb)
}
