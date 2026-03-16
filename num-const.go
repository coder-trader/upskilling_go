package main

import "fmt"

const Big = 1 << 2

func bishi() int {
	return Big

}

func ifelses() string {
	if v := bishi(); v > 2 {
		return fmt.Sprintf("bishi is big, %v\n", v)
	}
	return fmt.Sprintf("bishi is small, %v\n", bishi())
}
