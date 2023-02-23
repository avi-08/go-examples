package main

import "fmt"

func fibRecursive(n int, series *[]int) {
	if n < 0 {
		fmt.Println(fmt.Errorf("number of terms in the sequence must be >= 0. Provided: %d", n))
		return
	} else if n <= 1 {
		if n == 0 {
			return
		} else if n == 1 {
			*series = append(*series, 0)
			return
		}
	}
	if len(*series) < 2 {
		*series = append(*series, 0, 1)
	}
	if len(*series) >= n {
		return
	}
	p := len(*series)
	*series = append(*series, (*series)[p-1]+(*series)[p-2])
	fibRecursive(n, series)
}

func fibIterative(n int) []int {
	switch {
	case n < 0:
		fmt.Println(fmt.Errorf("number of terms in the sequence must be >= 0. Provided: %d", n))
		return []int{}
	case n == 0:
		return []int{}
	case n == 1:
		return []int{0}
	case n >= 2:
		series := make([]int, 0)
		series = append(series, 0, 1)
		for i := 2; i < n; i++ {
			series = append(series, series[i-1]+series[i-2])
		}
		return series
	}
	return []int{}
}

func main() {
	testTerms := []int{7, 0, 1, 2, -5, 10}
	for _, val := range testTerms {
		seq := fibIterative(val)
		seqR := make([]int, 0)
		fibRecursive(val, &seqR)
		fmt.Println("Fibonacci sequence till", val, "terms")
		fmt.Println("=> Iterative:", seq)
		fmt.Println("=> Recursive:", seqR)
	}
}
