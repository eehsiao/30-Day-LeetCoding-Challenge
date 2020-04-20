package main

func isHappy(n int) bool {
	var (
		runed = make(map[int]struct{}, 0)
		happy bool
	)

	for {
		if _, ok := runed[n]; ok {
			happy = false
			break
		}
		runed[n] = struct{}{}
		n = func(n int) (newN int) {
			for {
				newN += (n % 10) * (n % 10)
				if n /= 10; n == 0 {
					break
				}
			}
			return
		}(n)
		if n == 1 {
			happy = true
			break
		}
	}
	return happy
}
