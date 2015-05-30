package main

import "fmt"

type Quaternion int


func (x Quaternion) power(n int64) Quaternion {
	//
	// 1, -1 => 1, -1
	// 2, -2 => i, -i
	// 3, -3 => j, -j
	// 4, -4 => k, -k
	//

	table := [5][5]Quaternion {
		{0,  0,  0,  0,  0},
		{1,  1,  1,  1,  1},
		{1,  2, -1, -2,  1},
		{1,  3, -1, -3,  1},
		{1,  4, -1, -4,  1},
	}

	sign := 1

	if x < 0 {
		x = -x
		if n % 2 == 1 {
			sign = -sign
		}
	}

	if n > 4 {
		n %= 4
	}

	if sign == -1 {
		return -table[x][n]
	} else {
		return table[x][n]
	}
}

func (x Quaternion) multiply(y Quaternion) Quaternion {
	//
	// 1, -1 => 1, -1
	// 2, -2 => i, -i
	// 3, -3 => j, -j
	// 4, -4 => k, -k
	//

	table := [5][5]Quaternion {
		{0,  0,  0,  0,  0},
		{0,  1,  2,  3,  4},
		{0,  2, -1,  4, -3},
		{0,  3, -4, -1,  2},
		{0,  4,  3, -2, -1},
	}

	sign := 1
	if x < 0 {
		x = -x
		sign = -sign
	}

	if y < 0 {
		y = -y
		sign = -sign
	}

	if sign == -1 {
		return -table[x][y]
	} else {
		return table[x][y]
	}
}

func (x Quaternion) leftMultiply(y Quaternion) Quaternion {
	//
	// 1, -1 => 1, -1
	// 2, -2 => i, -i
	// 3, -3 => j, -j
	// 4, -4 => k, -k
	//

	table := [5][5]Quaternion {
		{0,  0,  0,  0,  0},
		{0,  1,  2,  3,  4},
		{0,  2, -1,  4, -3},
		{0,  3, -4, -1,  2},
		{0,  4,  3, -2, -1},
	}

	sign := 1
	if x < 0 {
		x = -x
		sign = -sign
	}

	if y < 0 {
		y = -y
		sign = -sign
	}

	if sign == -1 {
		return -table[y][x]
	} else {
		return table[y][x]
	}
}




func totalProduct(s string, repetition int64) Quaternion {
	var product Quaternion = 1

	for i := 0; i < len(s); i++ {
		product = product.multiply( Quaternion(s[i] - 'i') + 2)
	}
	return product.power(repetition)
}

func findFirstI(s string, repetition int64) int64 {
	var i int64
	var product Quaternion = 1

	for i = 0; i < int64(len(s)) * repetition && i < int64(len(s)) *4 ; i++ {
		product = product.multiply( Quaternion(s[i % int64(len(s))] - 'i') + 2)
		if product == 2 {
			return i
		}
	}
	return -1
}

func findLastK(s string, repetition int64) int64 {
	var i int64
	var product Quaternion = 1

	for i = 0; i < int64(len(s)) * repetition && i < int64(len(s)) *4 ; i++ {
		j := int64(len(s)) * repetition - i - 1
		product = product.leftMultiply( Quaternion(s[j % int64(len(s))] - 'i') + 2)
		if product == 4 {
			return j
		}
	}
	return -1	
}

func check(s string, repetition int64) bool {

	product := totalProduct(s, repetition)

	if product != -1 {
		return false
	}

	i := findFirstI(s, repetition)

	if i == -1 {
		return false
	}

	k := findLastK(s, repetition)

	if k == -1 {
		return false
	}
	return i < k
}

func main() {
	T := 0

	fmt.Scanf("%d\n", &T)

	for i := 1; i <= T; i++ {
		var X, L int64
		fmt.Scanf("%d %d\n", &X, &L)

		var input string
		fmt.Scanf("%s\n", &input)

		if check(input, L) {
			fmt.Printf("Case #%v: YES\n", i)
		} else {
			fmt.Printf("Case #%v: NO\n", i)
		}
	}
}