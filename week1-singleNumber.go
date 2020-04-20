package main

func singleNumber(nums []int) int {
	var (
		singleN   = make(map[int]struct{}, 0)
		multiN    = make(map[int]struct{}, 0)
		returnKey int
	)

	for _, n := range nums {
		if _, ok := multiN[n]; ok {
			continue
		}
		if _, ok := singleN[n]; ok {
			multiN[n] = struct{}{}
			delete(singleN, n)
		} else {
			singleN[n] = struct{}{}
		}
	}

	for k, _ := range singleN {
		returnKey = k
		break
	}

	return returnKey
}
