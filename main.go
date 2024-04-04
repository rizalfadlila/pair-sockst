package main

import "fmt"

func main() {
	cases := testCase()
	for i := 0; i < len(cases); i++ {
		fmt.Println(fmt.Sprintf("case %v, result : %v", i+1, sockMerchant(cases[i])))
	}
}

func sockMerchant(arr []int) int {
	pairs := map[int]int{}
	counter := 0
	for i := 0; i < len(arr); i++ {
		if _, ok := pairs[arr[i]]; ok {
			counter++
			delete(pairs, arr[i])
		}
		pairs[arr[i]]++
	}
	return counter
}

func testCase() [][]int {
	return [][]int{
		{10, 20, 20, 10, 10, 30, 50, 10, 20},
		//{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5},
		//{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5},
		//{10, 20, 30},
		//{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		//{1, 1, 3, 1, 2, 1, 3, 3, 3, 3},
	}
}
