package LRU

type ListNode struct {
	key, val  int
	pre, next *ListNode
}

type LRUCache struct {
	len, cap int
	hm       map[int]*ListNode
	head     *ListNode
	tail     *ListNode
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{}
	tail := &ListNode{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		cap:  capacity,
		head: head,
		tail: tail,
		hm:   make(map[int]*ListNode),
	}
}

func (this *LRUCache) Remove(node *ListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) Add(node *ListNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hm[key]; ok {
		// 删除链表所在位置
		this.Remove(v)
		// 移动到头指针
		this.Add(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hm[key]; ok {
		v.val = value
		// 删除链表所在位置
		this.Remove(v)
		// 移动到头指针
		this.Add(v)
	} else {
		head := &ListNode{key: key, val: value}
		if this.len >= this.cap {
			delete(this.hm, this.tail.pre.key)
			// 删除尾指针
			this.Remove(this.tail.pre)
		}
		this.len++
		this.hm[key] = head
		this.Add(head)
	}
}
