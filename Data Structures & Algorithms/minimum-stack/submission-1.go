type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
		}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// for 2 stack approach we compare while pushing itself 
	minVal := val
	if len(this.minStack) > 0 {
		if top := this.minStack[len(this.minStack)-1]; 
		top < val {
			minVal = top
		}
	}
		this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	// tmp := []int{}
	// min := this.Top()
	// // pop all values from stack and keep minimum 
	// for len(this.stack) > 0 {
	// 	val := this.stack[len(this.stack)-1]
	// 	this.stack = this.stack[:len(this.stack)-1]
	// 	min = findMin(min, val)
	// 	tmp = append(tmp, val)
	// }
	// // Push all values back into the stack
	// for i:=len(tmp)-1; i>=0; i--{
	// 	this.stack = append(this.stack, tmp[i])
	// } 
	// // for len(tmp) > 0 {
	// // 	val := tmp[len(tmp)-1]
	// // 	tmp = tmp[:len(tmp)-1]
	// // 	this.stack = append(this.stack, val)
	// // }

	// return min

	// Two stack approach
	return this.minStack[len(this.minStack)-1]
}

// func findMin(a,b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }