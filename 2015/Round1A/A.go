package main

import "fmt"

func check(m []int) (int, int) {
	method1, method2 := 0, 0
	minRate := -1
	for i := 0; i < len(m) - 1; i++ {
		if m[i+1] <= m[i] {
			method1 += m[i] - m[i+1]
			if m[i] - m[i+1] > minRate {
				minRate = m[i] - m[i+1]
			}
		}
	}

	for i := 0; i < len(m) - 1; i++ {
		if m[i] < minRate {
			method2 += m[i]
		} else {
			method2 += minRate
		}
	}


	return method1, method2
}

func main() {
	T := 0

	fmt.Scanf("%d\n", &T)

	for i := 1; i <= T; i++ {
		n := 0
		fmt.Scanf("%d\n", &n)

		m := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &m[j])
		}

		m1, m2 := check(m)
		fmt.Printf("Case #%v: %v %v\n", i, m1, m2)
	}

}