package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()

	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	if len(this.queue) != 0 {
		val := this.queue[len(this.queue)-1]
		this.queue = this.queue[:len(this.queue)-1]
		return val
	}

	return 0
}

func (this *MyStack) Top() int {
	if len(this.queue) != 0 {
		return this.queue[len(this.queue)-1]
	}
	return 0
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
