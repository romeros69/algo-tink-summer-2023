package main

import (
	"fmt"
	"sort"
)

func alg4() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	set := map[int]int{}
	curArr := make([]int, 0)
	for curLen := n; curLen > 1; curLen-- {
		if curLen == 2 {
			fmt.Println(curLen)
			return
		}
		for i := 0; i < curLen; i++ {
			set[a[i]]++
		}
		for _, val := range set {
			curArr = append(curArr, val)
		}
		sort.Ints(curArr)
		checkCountDif := 0
		for i := 0; i < len(curArr)-1; i++ {
			if curArr[i] != curArr[i+1] {
				checkCountDif++
			}
		}
		if checkCountDif > 1 {
			curArr = nil
			for k := range set {
				delete(set, k)
			}
			continue
		} else if checkCountDif == 0 {
			fmt.Println(curLen)
			return
		}
		for i := 0; i < len(curArr)-1; i++ {
			if curArr[i] != curArr[i+1] {
				if (i > 0) && (len(curArr)-(i+1) > 1) {
					continue
				} else {
					if (curArr[i+1] - curArr[i]) != 1 {
						continue
					} else {
						fmt.Println(curLen)
						return
					}
				}
			}
		}
		curArr = nil
		for k := range set {
			delete(set, k)
		}
	}
}
