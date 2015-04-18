package main

import "fmt"

func numberOfServices(m []int, t int) int {
	total := 0
	for i := 0; i < len(m); i++ {
		total += (t + m[i] - 1)/ m[i]
	}
	return total
}

func check(m []int, n int) int {
	minTime, maxTime := 0, 100000*1000000000

	for minTime + 1 < maxTime {
		t := minTime + (maxTime - minTime) / 2
		if numberOfServices(m, t) >= n {
			maxTime = t
		} else {
			minTime = t 
		}
	}

	//fmt.Println(minTime, numberOfServices(m, minTime), maxTime, numberOfServices(m, maxTime))

	remain := n - numberOfServices(m, minTime)
	for i := 0; i < len(m); i++ {
		if minTime % m[i] == 0 {
			remain--
			if remain == 0 {
				return i+1
			}
		}
	}
	panic("error")
}

func main() {
	T := 0

	fmt.Scanf("%d\n", &T)

	for i := 1; i <= T; i++ {
		b, n := 0, 0
		fmt.Scanf("%d%d\n", &b, &n)
		//fmt.Println(b, n)

		m := make([]int, b)
		for j := 0; j < b; j++ {
			fmt.Scanf("%d", &m[j])
			//fmt.Println(m[j])
		}
		//fmt.Scanf("\n")

		fmt.Printf("Case #%v: %v\n", i, check(m, n))
	}

}