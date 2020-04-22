package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func iif(l bool, a interface{}, b interface{}) interface{} {
	if l {
		return a
	}
	return b
}