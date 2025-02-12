package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte("aabbccc")
	// chars := []byte("aaabbaa")
	result_len := compress(chars)
	fmt.Printf("The length of the compressed string is: %d\n", result_len)
}

func compress_old(chars []byte) int {
	var last_char byte = chars[0]
	var start int
	var end int
	var compressed_str []byte
	num_ch := 1
	chars = append(chars, byte(0))
	for i := 1; i < len(chars); i++ {
		if chars[i] != last_char {
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
		last_char = chars[i]
	}
	chars = chars[:len(chars)-1]
	fmt.Println(string(chars))
	return len(chars)
}

func compressSlice(start int, end int, chars []byte, c byte) {
	for j := start; j < end; j++ {

	}
}

func compress(chars []byte) int {
	write := 0
	read := 0
	
	for read < len(chars) {
		char := chars[read]
		count := 0

		for read < len(chars) && chars[read] == char {
			read++
			count++
		}

		chars[write] = char
		write++
		
		if count > 1 {
			for _, digit_char := range strconv.Itoa(count){
				chars[write] = byte(digit_char)
				write++
			}
		}
	}
	fmt.Println(string(chars))
	return write
}