package min_stack

type MinStack struct {
	arr    []int
	minArr []int
}

func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {

	this.arr = append(this.arr, val)

	if len(this.minArr) == 0 {
		this.minArr = append(this.minArr, val)
		return
	}

	currMin := this.minArr[len(this.minArr)-1]

	if val < currMin {
		this.minArr = append(this.minArr, val)
	} else {
		this.minArr = append(this.minArr, currMin)
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]

}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}
