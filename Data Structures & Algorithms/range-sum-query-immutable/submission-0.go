type NumArray struct {
    arr []int
}


func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i:=1; i < len(nums); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
    return NumArray{
		arr: arr,
	}
}
func (this *NumArray) SumRange(left int, right int) int {
		// we need to subtract left-1 from right because the left already contains prefix sum  
		if left == 0 {
			return this.arr[right]
		}
    return this.arr[right] - this.arr[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */