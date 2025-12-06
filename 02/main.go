package main

import (
	"slices"
	"strconv"
	"strings"
)

func Run1(input []Range) int {
	acc := 0
	for _, r := range input {
		invalid := make(chan int)
		go get_invalid(r, 2, invalid)
		for v := range invalid {
			acc += v
		}
	}
	return acc
}

func Run2(input []Range) int {
	acc := 0
	for _, r := range input {
		max_reps := len(strconv.Itoa(r.last_id))
		seen := []int{}
		for reps := 2; reps <= max_reps; reps++ {
			invalid := make(chan int)
			go get_invalid(r, reps, invalid)
			for v := range invalid {
				if slices.Contains(seen, v) {
					continue
				}
				seen = append(seen, v)
				acc += v
			}
		}
	}
	return acc
}

func get_invalid(r Range, num_reps int, result chan int) {
	first_s := strconv.Itoa(r.first_id)
	last_s := strconv.Itoa(r.last_id)
	min_len := len(first_s)
	max_len := len(last_s)
	for cur_len := min_len; cur_len <= max_len; cur_len++ {
		if cur_len%num_reps != 0 {
			continue // cannot contain num_reps equal length sequences
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
	outer:
		for i := start_i; i <= end_i; i++ {
			s := strconv.Itoa(i)
			left := s[:cur_len/num_reps]
			for j := 1; j < num_reps; j++ {
				right := s[cur_len/num_reps*j : cur_len/num_reps*(j+1)]
				if left != right {
					continue outer
				}
			}
			result <- i
		}
	}
	close(result)
}
