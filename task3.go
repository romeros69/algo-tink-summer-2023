package main

import (
	"fmt"
	"strings"
)

func alg3() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	if n < 4 {
		fmt.Println(-1)
		return
	}
	curLenSub := 4
	l := 0
	r := 4
	res := 0
OuterLoop:
	for curLenSub <= n {
		for r != n+1 {
			if strings.Contains(s[l:r], "a") &&
				strings.Contains(s[l:r], "b") &&
				strings.Contains(s[l:r], "c") &&
				strings.Contains(s[l:r], "d") {
				res = curLenSub
				break OuterLoop
			}
			l++
			r++
		}
		curLenSub++
		l = 0
		r = curLenSub

	}
	if res != 0 {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}

}
