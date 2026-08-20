func topKFrequent(nums []int, k int) []int {
    // count := make(map[int]int)
    // for _, num := range nums {
    //     count[num]++
    // }
    // arr := make([][2]int, 0, len(count))
    // for i,v := range count{
    //     arr = append(arr, [2]int{i,v})
    // }
    // sort.Slice(arr, func(i, j int) bool {
    //     return arr[i][0] > arr[j][0]
    // })

    // res := make([]int, k)
    // for i:=0; i<k; i++{
    //     res[i]=arr[i][1]
    // }
    // return res

    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }
    arr := make([][]int, len(nums)+1)
    for num, freq := range count {
        arr[freq] = append(arr[freq],num)
    }
    result := []int{}

    for i:= len(arr)-1; i>=0 && len(result) < k; i-- {
        if arr[i] != nil {
            result = append(result, arr[i]...)
        }
    }
    return result
 }
