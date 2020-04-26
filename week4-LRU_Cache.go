// Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
// The cache is initialized with a positive capacity.
// Follow up:
// Could you do both operations in O(1) time complexity?
// Example:
// LRUCache cache = new LRUCache( 2 /* capacity */ );
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
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

package main

type dataNode struct {
	k, v              int
	preNode, nextNode *dataNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*dataNode
	first    *dataNode
	last     *dataNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*dataNode),
	}
}

func (this *LRUCache) Get(k int) int {
	if node, ok := this.cache[k]; ok {
		this.moveToFirst(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(k int, v int) {
	var (
		node *dataNode
		ok   bool
	)
	if node, ok = this.cache[k]; !ok {
		if len(this.cache) >= this.capacity {
			this.removeLast()
		}
		node = &dataNode{}
	}
	node.k = k
	node.v = v

	this.moveToFirst(node)
	this.cache[k] = node
}

func (this *LRUCache) removeLast() {
	if this.last != nil {
		last := this.last
		if last.preNode != nil {
			this.last.preNode.nextNode = nil
			this.last = this.last.preNode
		} else {
			this.first = nil
			this.last = nil
		}
		delete(this.cache, last.k)
	}
}

func (this *LRUCache) moveToFirst(node *dataNode) {
	if node == this.first {
		return
	}
	if node.preNode != nil {
		node.preNode.nextNode = node.nextNode
	}
	if node.nextNode != nil {
		node.nextNode.preNode = node.preNode
	}
	if this.last == node {
		this.last = node.preNode
	}
	if this.first != nil {
		node.nextNode = this.first
		this.first.preNode = node
	}
	this.first = node
	node.preNode = nil
	if this.last == nil {
		this.last = this.first
	}
}
