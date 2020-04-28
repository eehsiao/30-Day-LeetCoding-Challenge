// Author :		Eric<eehsiao@gmail.com>
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/

package main

func main() {
	// cache := Constructor(2 /* capacity */)

	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// fmt.Println(cache.Get(1)) // returns 1
	// cache.Put(3, 3)           // evicts key 2
	// fmt.Println(cache.Get(2)) // returns -1 (not found)
	// cache.Put(4, 4)           // evicts key 1
	// fmt.Println(cache.Get(1)) // returns -1 (not found)
	// fmt.Println(cache.Get(3)) // returns 3
	// fmt.Println(cache.Get(4)) // returns 4

	// firstUnique := FirstUnique_Constructor([]int{7, 7, 7, 7, 7, 7})
	// firstUnique.prt()
	// firstUnique.ShowFirstUnique() // return -1
	// firstUnique.Add(7)            // the queue is now [7,7,7,7,7,7,7]
	// firstUnique.prt()
	// firstUnique.Add(3) // the queue is now [7,7,7,7,7,7,7,3]
	// firstUnique.prt()
	// firstUnique.Add(3) // the queue is now [7,7,7,7,7,7,7,3,3]
	// firstUnique.prt()
	// firstUnique.Add(7) // the queue is now [7,7,7,7,7,7,7,3,3,7]
	// firstUnique.prt()
	// firstUnique.Add(17) // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
	// firstUnique.prt()
	// firstUnique.ShowFirstUnique() // return 17

	firstUnique := FirstUnique_Constructor([]int{2, 3, 5})
	firstUnique.prt()
	firstUnique.ShowFirstUnique() // return -1
	firstUnique.Add(5)            // the queue is now [7,7,7,7,7,7,7]
	firstUnique.prt()
	firstUnique.Add(2) // the queue is now [7,7,7,7,7,7,7,3]
	firstUnique.prt()
	firstUnique.Add(3) // the queue is now [7,7,7,7,7,7,7,3,3]
	firstUnique.prt()
	firstUnique.ShowFirstUnique() // return 17
}
