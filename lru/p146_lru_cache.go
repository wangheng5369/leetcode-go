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
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.moveToHead(node)
	} else {
		node := &Node{key: key, value: value}
		c.cache[key] = node
		if len(c.cache) > c.capacity {
			c.removeTail()
			delete(c.cache, c.tail.prev.key)
		}
	}
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) addToHead(node *Node) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeTail() *Node {
	node := c.tail.prev
	c.removeNode(node)
	return node
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
