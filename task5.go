package main

import "fmt"

var mapa = make(map[int]bool)

func checkSum(a []int) bool {
	s := 0
	for _, v := range a {
		s += v
	}
	if s == 0 {
		return true
	}
	return false
}

func checkInMap(i, j int) bool {
	key := i*10 + j
	if _, ok := mapa[key]; ok {
		return true
	}
	return false
}

func addToMap(i, j int) {
	key := i*10 + j
	mapa[key] = true
}

// i - начало отрезка
// j - конец отрезка (включительно)
func countOtrezkov(gen []int, i, j int) int {
	count := 0
	for p := j; p < len(gen); p++ {
		if !checkInMap(i, p) {
			count++
			addToMap(i, p)
		}
	}
	return count
}

func alg5() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var cur []int
	result := 0 // суммарное количество нормльных отрезков
	i := 0
	j := 2 // в данном случае не включительно
	for {
		if i == (len(a) - 1) {
			break
		}
		cur = a[i:j]
		if checkSum(cur) {
			result = result + countOtrezkov(a, i, j-1)
		}
		if j == len(a) {
			i++
			j = i + 2
		} else {
			j++
		}
	}
	fmt.Println(result)
}
