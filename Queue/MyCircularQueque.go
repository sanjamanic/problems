package queue

type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	var this MyCircularQueue
	this.size = k
	this.data = make([]int, k)
	this.head = -1
	this.tail = -1
	return this
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsEmpty() {
		this.tail = 0
		this.head = 0
		this.data[this.tail] = value
		return true
	} else {
		if this.IsFull() {
			return false
		} else {
			this.tail++
			if this.tail == this.size {
				this.tail = 0
			}
			this.data[this.tail] = value
			return true
		}
	}
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.head++
		if this.head == this.tail+1 {
			this.tail = -1
			this.head = -1
		}
		if this.head == this.size {
			this.head = 0
		}
		return true
	}
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() {
		return this.data[this.head]
	} else {
		return -1
	}
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.data[this.tail]
	} else {
		return -1
	}
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if (this.tail == -1) && (this.head == -1) {
		return true
	} else {
		return false
	}

}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.head == (this.tail+1)%this.size {
		return true
	} else {
		return false
	}
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
