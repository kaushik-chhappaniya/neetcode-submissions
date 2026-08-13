func isAnagram(s string, t string) bool {
    if len(s) != len(t){ 
        return false
    }

    sL := strings.ToLower(s)
    tL := strings.ToLower(t)

    var arr = make([]int,26)
    for i := range(s) {
        arr[sL[i]-'a']++
        arr[tL[i]-'a']--
    }
    for _, v:= range arr {
        if v != 0 {
            return false
        }
    }
    return true

    // var hm = make(map[rune]int)
    // for _,v := range(s) {
    // hm[v]++
    // }
    // for _, v:=range(t) {
    //     if _, ok:= hm[v]; ok && hm[v] > 0{
    //         hm[v]--
    //     } else {
    //         return false
    //     } 
    // }
    // return true
}
