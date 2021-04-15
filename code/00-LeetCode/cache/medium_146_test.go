package cache

/*
	https://leetcode-cn.com/problems/lru-cache/

	LRU 缓存淘汰算法实现

	Least Recently Used,我们认为最近使用过的数据应该是有用的,很久没有使用的数据应该是无用的,
	内存满了就先删掉那些很久没用过的数据

	思路：
		- 核心数据结构——哈希链表 LinkedHashMap
		- 首先借助链表的有序性使得链表元素维持插入顺序
		- 借助哈希映射的快速访问使得我们可以通过O(1)的时间复杂度访问链表中的任意元素

	分析
		1. 如果每次默认从链表尾部添加元素,那么显然越靠尾部的元素就越是最近使用的,越靠近头部的元素就是越久未使用的
		2. 对于某一个key,可以通过哈希表快速定位到链表中的节点,从而取得对应得Val
		3. 链表显然是支持在任意位置快速插入和删除的,改改指针就行。
		只不过传统的链表无法按照索引快速访问某一个位置的元素,而这里借助哈希表,可以通过key快速映射到任意一个链表节点,
		然后进行插入和删除
*/

type LRUCache struct {
	// Cap 最大容量
	Cap int

	// Data 存储节点数据
	Data map[int]*Node

	// Cache 双向链表：任意位置快速插入和删除,而删除操作不仅要得到该节点本身的指针,
	// 也要操作其其前驱节点的指针,双向链表支持直接查找前驱
	Cache *DuLinkedList
}

type Node struct {
	// key ,value
	key, val int

	// prev 前驱指针；next 后继指针
	prev, next *Node
}

func newNode(key, val int) *Node {
	return &Node{key: key, val: val}
}

// 双向链表数据结构
type DuLinkedList struct {
	// size 链表元素数量
	size int
	// 头尾虚节点
	head, tail *Node
}

// newDLinkedNode 初识化双向链表的数据
func newDuLinkedList() *DuLinkedList {
	dl := new(DuLinkedList)
	dl.head = newNode(0, 0)
	dl.tail = newNode(0, 0)

	// 双向链表的空节点定义
	dl.head.next = dl.tail // 头指尾
	dl.tail.prev = dl.head // 尾指头
	dl.size = 0

	return dl
}

// Constructor 以正整数作为容量 capacity 初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Data:  map[int]*Node{},
		Cache: newDuLinkedList(),
	}
}

// Get 如果关键字 key 存在于缓存中,则返回关键字的值,否则返回 -1
func (lru *LRUCache) Get(key int) int {
	// 先查哈希表中是否有数据，没有则返回-1
	if _, ok := lru.Data[key]; ok {
		// 将该数据提升为最近使用的
		lru.makeRecently(key)
		return lru.Data[key].val
	}
	return -1
}

// 如果关键字已经存在则变更其数据值；如果关键字不存在,则插入该组「关键字-值」。
// 当缓存容量达到上限时,它应该在写入新数据之前删除最久未使用的数据值,从而为新的数据值留出空间。
func (lru *LRUCache) Put(key int, value int) {
	// 先查哈希表中是否有数据，
	if _, ok := lru.Data[key]; ok {
		// 1. 从哈希表中删除旧的数据
		lru.deleteKey(key)

		// 2.新插入的数据为最近使用的数据
		lru.addRecently(key, value)
		return
	}
	// 如果LRUCache 容量已满，删除最久未被使用的数据
	if lru.Cap == lru.Cache.Size() {
		lru.removeNotUsed()
	}

	lru.addRecently(key, value)
}

// -------------begin LRUCache辅助操作-----------
// 先不着急实现LRU算法的get和put方法。由于要同时维护一个双链表的cache和一个哈希表，很容易漏掉一些操作
// 比如删除某个key时，在cache中删除了对应的node，但是却忘记在哈希表中删除key
// 解决这种问题的有效方法：在两种数据结构之上提供一层抽象API,尽量让LRU的主方法get和put避免直接操作cache和哈希表
// 的细节，故实现以下辅助函数

// makeRecently 将某个key提升为最近使用的
func (lru *LRUCache) makeRecently(key int) {
	x := lru.Data[key]
	// 先从链表中删除这个节点
	lru.Cache.remove(x)
	// 重新插到链表尾
	lru.Cache.addLast(x)
}

// addRecently 添加最近使用的元素
func (lru *LRUCache) addRecently(key, val int) {
	x := newNode(key, val)
	// 链表尾部就是最近使用的元素
	lru.Cache.addLast(x)
	// 别忘了在哈希表中添加key的映射
	lru.Data[key] = x
}

// deleteKey 删除某一个key
func (lru *LRUCache) deleteKey(key int) {
	x := lru.Data[key]
	// 从链表中删除
	lru.Cache.remove(x)

	// 从哈希表中删除
	delete(lru.Data, key)
}

// removeNotUsed 删除最久未使用的元素，其实就是链表中的第一个元素
func (lru *LRUCache) removeNotUsed() {
	delete(lru.Data, lru.Cache.removeFirst().key)
}

// -------------begin LRUCache辅助操作-----------

// -------------begin 双向链表操作-----------

// addLast 在链表尾部添加节点,时间复杂度度O(1)
func (dl *DuLinkedList) addLast(x *Node) {
	x.prev = dl.tail.prev
	x.next = dl.tail

	dl.tail.prev.next = x
	dl.tail.prev = x
	dl.size++
}

// remove 删除链表中的x节点(x一定存在), 时间复杂度O(1)
func (dl *DuLinkedList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	dl.size--
}

// removeFirst 删除链表中的第一个节点,并返回该节点, 时间复杂度O(1)
func (dl *DuLinkedList) removeFirst() *Node {
	if dl.head.next == dl.tail {
		return nil
	}

	first := dl.head.next // 虚拟头节点的next才是真正的头节点
	dl.remove(first)
	return first
}
func (dl *DuLinkedList) Size() int {
	return dl.size
}

// -------------end 双向链表操作-----------
