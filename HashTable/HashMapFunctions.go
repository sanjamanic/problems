package hashtable

type MyHashMap struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	var table MyHashMap
	table.m = make(map[int]int)
	return table
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
	return
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	var ok bool
	ok = this.Contains(key)
	if ok {
		return this.m[key]
	} else {
		return -1
	}

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	var ok bool
	ok = this.Contains(key)
	if ok {
		delete(this.m, key)
	}
	return
}

/** Returns true if this set contains the specified element */
func (this *MyHashMap) Contains(key int) bool {
	var ok bool
	_, ok = this.m[key]
	return ok
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
