package queue

func numIslands(grid [][]byte) int {
	var this MyCircularQueue
	var ic, jc, m, n, l int
	cnt := 0
	m = len(grid)
	if m > 0 {
		n = len(grid[0])
		l = 2 * m * n
		this.size = l
		this.data = make([]int, l)
		this.head = -1
		this.tail = -1
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == '1' {
					cnt = cnt + 1
					this.EnQueue(i)
					this.EnQueue(j)
					grid[i][j] = '0'
					for {
						ic = this.Front()
						this.DeQueue()
						jc = this.Front()
						this.DeQueue()
						this.AddIndeces(ic, jc, m, n, grid)
						if this.IsEmpty() {
							break
						}
					}
				}
			}
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

func (this *MyCircularQueue) AddIndeces(i int, j int, m int, n int, grid [][]byte) {

	for k := -1; k < 2; k++ {
		for l := -1; l < 2; l++ {
			ic := i + k
			jc := j + l
			if ic != i && jc != j {
				continue
			}
			if ic >= m || ic < 0 || jc >= n || jc < 0 {
				continue
			}
			if grid[ic][jc] == '1' {
				this.EnQueue(ic)
				this.EnQueue(jc)
				grid[ic][jc] = '0'
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
