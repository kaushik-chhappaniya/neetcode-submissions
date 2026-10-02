func canPlaceFlowers(flowerbed []int, n int) bool {
    emp:=0
    if flowerbed[0] == 0 {
        emp = 1
    }
    for _, v := range flowerbed {
        if v == 1 {
            n -= (emp -1)/2
            emp = 0
        } else {
            emp++
        }
    }
    n -= emp / 2
    return n <= 0
}