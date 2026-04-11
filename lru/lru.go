package lru

import "fmt"

type Node struct {
	Key   any
	Value any
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[any]*Node
	head     *Node
	tail     *Node
}

func New(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	h := &Node{}
	t := &Node{}
	h.Next = t
	t.Prev = h
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[any]*Node, capacity),
		head:     h,
		tail:     t,
	}
}

func (c *LRUCache) Get(key any) (any, error) {
	node, ok := c.cache[key]
	if !ok {
		return nil, fmt.Errorf("key %v not found", key)
	}
	c.moveToHead(node)
	return node.Value, nil
}

func (c *LRUCache) Put(key, value any) {
	if node, ok := c.cache[key]; ok {
		node.Value = value
		c.moveToHead(node)
		return
	}
	node := &Node{Key: key, Value: value}
	c.cache[key] = node
	c.addToHead(node)
	if len(c.cache) > c.capacity {
		removed := c.removeTail()
		delete(c.cache, removed.Key)
	}
}

func (c *LRUCache) addToHead(node *Node) {
	node.Prev = c.head
	node.Next = c.head.Next
	c.head.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *Node {
	node := c.tail.Prev
	c.removeNode(node)
	return node
}

func (c *LRUCache) Len() int {
	return len(c.cache)
}
