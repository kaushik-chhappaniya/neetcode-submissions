type StockSpanner struct {
	arr []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		arr: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.arr = append(this.arr, price)
	i := len(this.arr) - 2
	for i>= 0 && this.arr[i] <= price {
		i--
	}
	return len(this.arr) - i - 1
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 