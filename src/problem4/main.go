package main

func sum_to_n_a(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_b(n int) int {
	if n == 1 {
		return n
	}
	return n + sum_to_n_b(n-1)
}

func sum_to_n_c(n int) int {
	return n * (n + 1) / 2
}
