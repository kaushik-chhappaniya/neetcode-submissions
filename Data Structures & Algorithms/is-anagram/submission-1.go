func isAnagram(s string, t string) bool {
    var hm = make(map[rune]int)
    if len(s) != len(t){ return false}
    for _,v := range(s) {
    hm[v]++
    }
    for _, v:=range(t) {
        if _, ok:= hm[v]; ok && hm[v] > 0{
            hm[v]--
        } else {
            return false
        } 
    }
    return true
}
