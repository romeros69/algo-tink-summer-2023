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
	const (
		T0 = iota // не имеет типа
		T1        // возрастание
		T2        // убывание
	)
	upType := T0
	prev := numSlice[0]
	for i := 1; i < 4; i++ {
		if upType == 0 {
			if prev < numSlice[i] {
				upType = T1
			} else if prev > numSlice[i] {
				upType = T2
			}
		} else {
			if (upType == T1 && prev > numSlice[i]) || (upType == T2 && prev < numSlice[i]) {
				res = false
				break
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
