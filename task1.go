package main

import (
	"fmt"
)

func alg1() {
	numSlice := [4]int{}
	fmt.Scanf("%d %d %d %d",
		&numSlice[0],
		&numSlice[1],
		&numSlice[2],
		&numSlice[3],
	)
	res := true
	upType := 0 // 1 по возрастанию, 2 по убыанию
	prev := numSlice[0]
	for i := 1; i < 4; i++ {
		if upType == 0 {
			if prev < numSlice[i] {
				upType = 1
			} else if prev > numSlice[i] {
				upType = 2
			}
		} else {
			if upType == 1 {
				if prev > numSlice[i] {
					res = false
					break
				}
			} else {
				if prev < numSlice[i] {
					res = false
					break
				}
			}
		}
		prev = numSlice[i]
	}
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
