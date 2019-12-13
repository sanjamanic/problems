package queue

import "fmt"

func openLock(deadends []string, target string) int {
	var current []int
	var targetInt []int
	var deadendsInt [][]int
	var cnt int = 0
	var sz int = 5
	var str string
	var this MyCircularQueue
	var flag bool = false

	current = make([]int, sz)
	l := len(deadends)
	deadendsInt = make([][]int, l)
	targetInt = make([]int, sz-1)
	for i := 0; i < l; i++ {
		str = deadends[i]
		targetInt = StringToInt(str)
		deadendsInt[i] = make([]int, sz-1)
		deadendsInt[i] = targetInt
	}

	targetInt = StringToInt(target)

	l = 10 * 10 * 10 * 10 * 5
	this.size = l
	this.data = make([]int, l)
	this.head = -1
	this.tail = -1

	for i := 0; i < sz; i++ {
		this.EnQueue(current[i])
	}
	for {
		for i := 0; i < sz; i++ {
			current[i] = this.Front()
			this.DeQueue()
		}
		//fmt.Println(current)
		flag = true
		for i := 1; i < sz; i++ {
			if current[i] != targetInt[i-1] {
				flag = false
				break
			}
		}
		if flag {
			cnt = current[0]
			break
		}

		flag = this.CheckDeadends(current[1:sz], deadendsInt)
		if !flag {
			deadendsInt = append(deadendsInt, current[1:sz])
		}
		this.AddChildrenToQueue(current, deadendsInt)
		if this.IsEmpty() {
			cnt = -1
			break
		}
	}
	return cnt
}

type MyCircularQueue struct {
	data []int
	head int
	tail int
	size int
}

func StringToInt(str string) []int {
	l := len(str)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = int(str[i]) - 48
	}
	return res
}

func (this *MyCircularQueue) CheckDeadends(current []int, deadends [][]int) bool {
	l := len(deadends)
	sz := len(deadends[0])
	flag := false
	for k := 0; k < l; k++ {
		flagInt := true
		for j := 0; j < sz; j++ {
			if current[j] != deadends[k][j] {
				flagInt = false
				break
			}
		}
		if flagInt {
			flag = true
			break
		}
	}
	return flag
}

func (this *MyCircularQueue) AddChildrenToQueue(current []int, deadends [][]int) {
	hlp := [8][4]int{{1, 0, 0, 0}, {-1, 0, 0, 0}, {0, 1, 0, 0}, {0, -1, 0, 0}, {0, 0, 1, 0}, {0, 0, -1, 0}, {0, 0, 0, 1}, {0, 0, 0, -1}}
	sz := len(current)
	child := make([]int, sz)
	child[0] = current[0] + 1

	for i := 0; i < 8; i++ {
		for j := 1; j < sz; j++ {
			child[j] = current[j] + hlp[i][j-1]
			if child[j] > 9 {
				child[j] = 0
			}
			if child[j] < 0 {
				child[j] = 9
			}
		}
		flag := this.CheckDeadends(child[1:sz], deadends)
		if !flag {
			for j := 0; j < sz; j++ {
				this.EnQueue(child[j])
			}
		}
	}
	return
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
		return -1 //this.data[0]
	}
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.data[this.tail]
	} else {
		return -1 //this.data[0]
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
