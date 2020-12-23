package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	// PINALIGYAIPPPI%%
	// PINALSIGYAHRPI
	// convert("PAYPALISHIRING", 4)
	// 1    7
	// 2  6 8
	// 3 5  9
	// 4
	convert("123456789", 4)
}

// convert
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	l := len(s)
	// every group digit num is (numRows - 1) * 2
	space := (numRows - 1) * 2
	fmt.Printf("space: %d\n", space)
	// loop numbers
	loopNum := int(math.Ceil(float64(l) / float64(space)))
	fmt.Printf("space: %d length: %d\n", space, l)
	fmt.Printf("%d\n", loopNum)
	var buf bytes.Buffer
	for i := 0; i < numRows; i++ {
		num := (numRows - i - 1) * 2
		for j := 0; j < loopNum && i+(j*space) < l; j++ {
			fmt.Printf("%c", s[i+(j*space)])
			buf.WriteByte(s[i+(j*space)])
			if i != 0 && i != numRows-1 && i+(j*space)+num < l {
				fmt.Printf("%c", s[i+(j*space)+num])
				buf.WriteByte(s[i+(j*space)+num])
			}
		}
	}
	return buf.String()
}
