package main

import (
	"strconv"
	"strings"
)

func Run(input []Range) int {
	acc := 0
	for _, r := range input {
		invalid := make(chan int)
		go get_invalid(r, invalid)
		for v := range invalid {
			acc += v
		}
	}
	return acc
}

func get_invalid(r Range, result chan int) {
	first_s := strconv.Itoa(r.first_id)
	last_s := strconv.Itoa(r.last_id)
	min_len := len(first_s)
	max_len := len(last_s)
	for cur_len := min_len; cur_len <= max_len; cur_len++ {
		if cur_len%2 != 0 {
			continue // cannot contain same equal length sequences
		}

		var start, end string
		if cur_len == len(first_s) {
			start = first_s
		} else { // cur_len > len(first_s)
			start = "1" + strings.Repeat("0", cur_len-1)
		}
		if cur_len == len(last_s) {
			end = last_s
		} else { // cur_len < len(last_s)
			end = strings.Repeat("9", cur_len)
		}

		// yeah, this can be optimized further, too lazy right now
		start_i, _ := strconv.Atoi(start)
		end_i, _ := strconv.Atoi(end)
		for i := start_i; i <= end_i; i++ {
			s := strconv.Itoa(i)
			left := s[cur_len/2:]
			right := s[:cur_len/2]
			if left == right {
				result <- i
			}
		}
	}
	close(result)
}
