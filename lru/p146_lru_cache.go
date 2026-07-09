package lru

// p146 LRU 缓存
// 请你设计并实现一个满足 LRU（最近最少使用）缓存约束的数据结构。
//
// LRUCache 类：
// - LRUCache(int capacity) 用正整数 capacity 初始化 LRU 缓存
// - int get(int key) 如果 key 存在于缓存中，则获取该 key 对应的 value，否则返回 -1
// - void put(int key, int value) 如果 key 已存在，则更新其对应的 value；
//   如果 key 不存在，则插入该 key-value 对。
//   当缓存容量达到 capacity 时，应该在插入新元素之前，
//   删除最近最少使用的元素（即最长时间未被访问的元素）。
//
// 你必须在 O(1) 的时间复杂度内完成 get 和 put 操作。
//
// 示例 1：
// 输入：
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出：[null, null, null, 1, null, 2, null, -1, 3, 4]
// 解释：
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存变成 {1=1}
// lRUCache.put(2, 2); // 缓存变成 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1，缓存变成 {2=2, 1=1}
// lRUCache.put(3, 3); // 该操作会使 key=2 失效，缓存变成 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1（key=2 已失效）
// lRUCache.put(4, 4); // 该操作会使 key=1 失效，缓存变成 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1（key=1 已失效）
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
// 提示：
// - 1 <= capacity <= 3000
// - 0 <= key <= 10^4
// - 0 <= value <= 10^5
// - 最多调用 get 和 put 10^5 次

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node = &Node{key: key, value: value}
		this.addToHead(node)
		this.cache[key] = node
		if len(this.cache) > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
		}
	}
}

func (this *LRUCache) addToHead(node *Node) {
	t := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = t
	t.prev = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	removed := this.tail.prev
	this.removeNode(removed)
	delete(this.cache, removed.key)
	return removed
}

func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}
