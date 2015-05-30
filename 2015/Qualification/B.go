package main

import "fmt"

func check(p []int, sliceMax int) int {

	step := 0
	realMax := 0

	for i := 0; i < len(p); i++ {
		if p[i] > sliceMax {
			step += (p[i] - 1) / sliceMax
			realMax = sliceMax
		} else if p[i] > realMax {
			realMax = p[i]
		}
	}
	return step + realMax
}

func main() {
	T := 0

	fmt.Scanf("%d\n", &T)

	for i := 1; i <= T; i++ {
		D := 0
		fmt.Scanf("%d\n", &D)

		max := 0
		P := make([]int, D)
		for j := 0; j < D; j++ {
			fmt.Scanf("%d", &P[j])

			if P[j] > max {
				max = P[j]
			}
		}

		min := check(P, max)
		for j := 1; j < max; j++ {
			step := check(P, j)
			if step < min {
				min = step
			}
		}

		fmt.Printf("Case #%v: %v\n", i, min)
	}
}