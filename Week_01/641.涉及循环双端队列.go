package Week_01

//设计实现双端队列。
//你的实现需要支持以下操作：
//
//
// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。
//
//
// 示例：
//
// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
// 
//
//
//
// 提示：
//
//
// 所有值的范围为 [1, 1000]
// 操作次数的范围为 [1, 1000]
// 请不要使用内置的双端队列库。
//
// Related Topics 设计 队列

type MyCircularDeque struct {
	front *node
	last  *node
	size  int
	cap   int
}

type node struct {
	Val  int
	Prev *node
	Next *node
}

/** 初始化双端队列. 设置队列长度为 k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{cap: k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size >= this.cap { //队列已满
		return false
	}
	n := &node{
		Val: value,
	}

	if this.size == 0 { //双端队列为空
		this.front = n
		this.last = n
		this.size++
		return true
	} else if this.size < this.cap { //双端队列未满
		n.Next = this.front
		this.front.Prev = n
		this.front = n
		this.size++
		return true
	}
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size >= this.cap { //队列已满
		return false
	}

	n := &node{
		Val: value,
	}

	if this.size == 0 { //双端队列为空
		this.front = n
		this.last = n
		this.size++
		return true
	} else if this.size < this.cap { //双端队列未满
		n.Prev = this.last
		this.last.Next = n
		this.last = n
		this.size++
		return true
	}
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 { //队列已空
		return false
	}

	if this.size == 1 { //只剩一个元素
		this.front = nil
		this.last = nil
		this.size = 0
		return true
	} else { //队列还有多个元素
		this.front = this.front.Next
		this.front.Prev = nil
		this.size--
		return true
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}

	if this.size == 1 {
		this.front = nil
		this.last = nil
		this.size = 0
		return true
	} else {
		this.last = this.last.Prev
		this.last.Next = nil
		this.size--
		return true
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}

	return this.front.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}

	return this.last.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
