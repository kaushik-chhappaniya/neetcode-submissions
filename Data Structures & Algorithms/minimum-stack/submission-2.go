type MinStack struct{ 
	minVal int
	nums []int
}

func Constructor () MinStack {
	return MinStack {
		minVal: math.MaxInt64,
		nums: []int{},
	}
}

func (this *MinStack) Push (val int) {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, 0)
		this.minVal = val
	} else {
		this.nums = append(this.nums, val - this.minVal)
		if val < this.minVal { 
			this.minVal = val
		}
	}
}

func (this *MinStack) Pop () {
	if len(this.nums) == 0 {
		return 
	}
	pop := this.nums[len(this.nums)-1]
	if pop < 0 {
		this.minVal = this.minVal - pop
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	top := this.nums[len(this.nums)-1]
	if top > 0 {
		return top + this.minVal
	}
	return this.minVal
}

func (this *MinStack) GetMin() int {
	return this.minVal
}