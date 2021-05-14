package main

type MyQueue struct {
	s, e []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if len(this.s) > 0 {
		this.e = append(this.e, this.s...)
		this.s = this.s[:0]
	}
	this.s = append(this.s, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.s) > 0 {
		this.e = append(this.e, this.s...)
		this.s = this.s[:0]
	}
	val := this.e[0]
	this.e = this.e[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s) > 0 {
		this.e = append(this.e, this.s...)
		this.s = this.s[:0]
	}
    return this.e[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s) == 0 && len(this.e) == 0
}
