package main

import (
	"fmt"
	"strconv"
)

func main() {
	// chars := []byte("aabbccc")
	chars := []byte("aaabbaa")
	result_len := compress(chars)
	fmt.Printf("The length of the compressed string is: %d\n", result_len)
	fmt.Println(byte('b'))
}

func compress(chars []byte) int {
	var last_char byte
	var start int
	var end int
	var compressed_str []byte
	num_ch := 1
	chars = append(chars, byte(0))
	for i, ch := range chars {
		if ch != last_char {
			switch num_ch {
			case 1:
				// Do nothing because non-repeated char cannot be compressed with this algorithm
				;
			default:
				fmt.Println("Foo, i:", i)
				start = i - num_ch
				end = i
				num_ch_bytes := []byte(strconv.Itoa(num_ch))
				compressed_str = append([]byte{last_char}, num_ch_bytes...)
				chars = append(chars[:start], append(compressed_str, chars[end:]...)...)
			}
			num_ch = 1
		} else {
			num_ch++
		}
		last_char = ch
	}
	chars = chars[:len(chars)-1]
	fmt.Println(string(chars))
	return len(chars)
}