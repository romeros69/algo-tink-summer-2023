package main

import "fmt"

func alg2() {
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)
	allCount := n * k // всего проверок
	lol := allCount / m
	if allCount%m == 0 {
		fmt.Println(lol)
	} else {
		fmt.Println(lol + 1)
	}
}
