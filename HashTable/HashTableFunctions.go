package hashtable

type MyHashSet struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	var table MyHashSet
	table.m = make(map[int]int)
	return table
}

func (this *MyHashSet) Add(key int) {
	var ok bool
	ok = this.Contains(key)
	if !ok {
		this.m[key] = 0
	}
	return

}

func (this *MyHashSet) Remove(key int) {
	var ok bool
	ok = this.Contains(key)
	if ok {
		delete(this.m, key)
	}
	return
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	var ok bool
	_, ok = this.m[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
