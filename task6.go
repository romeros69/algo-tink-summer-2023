package main

import (
	"fmt"
	"sort"
)

type node struct {
	min, max, cur int
}

// текущая сумма оценок
func getSum(mar []node) int {
	sum := 0
	for _, v := range mar {
		sum += v.cur
	}
	return sum
}

func alg6() {
	var n, sMax int
	fmt.Scan(&n)
	fmt.Scan(&sMax)
	mar := make([]node, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &mar[i].min, &mar[i].max)
		mar[i].cur = mar[i].max
	}
	sort.Slice(mar, func(i, j int) bool {
		if mar[i].max < mar[j].max {
			return true
		} else if mar[i].max > mar[j].max {
			return false
		}
		return mar[i].min < mar[j].min
	})
	if getSum(mar) > sMax {
	OuterLoop:
		for i := 0; i < n; i++ {
			for mar[i].cur > mar[i].min {
				if mar[i].cur == 12 {
					break
				}
				mar[i].cur--
				if getSum(mar) <= sMax {
					break OuterLoop
				}
			}

		}
		fmt.Println(mar[n/2].cur)
	} else {
		fmt.Println(mar[n/2].cur)
	}
}
