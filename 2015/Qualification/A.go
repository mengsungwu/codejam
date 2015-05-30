package main

import "fmt"

func check(s string) int {
	total := int(s[0]) - '0'
	helper := 0
	for i := 1; i < len(s); i++ {
		if total < i {
			helper += (i - total)
			total += (i - total)
		}

		total += int(s[i] - '0')
	}
	return helper
}

func main() {
	T := 0

	fmt.Scanf("%d\n", &T)

	for i := 1; i <= T; i++ {
		maxS := 0
		digits := ""
		fmt.Scanf("%d %s\n", &maxS, &digits)


		fmt.Printf("Case #%v: %v\n", i, check(digits))
	}

}