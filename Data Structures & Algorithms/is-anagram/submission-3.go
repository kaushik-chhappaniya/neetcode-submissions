func isAnagram(s string, t string) bool {
    if len(s) != len(t){ 
        return false
    }

    sL := strings.ToLower(s)
    tL := strings.ToLower(t)

    // FULL OPTIMISED
    mapS, mapT := make(map[rune]int), make(map[rune]int)
    for i, v:=range(sL) {
        mapS[v]++
        mapT[rune(tL[i])]++
    }

    for i,v := range mapS {
        if v != mapT[i] {
            return false
        }
    }
    return true

    // MEDIUM OPTIMISED
    // var arr = make([]int,26)
    // for i := range(s) {
    //     arr[sL[i]-'a']++
    //     arr[tL[i]-'a']--
    // }
    // for _, v:= range arr {
    //     if v != 0 {
    //         return false
    //     }
    // }
    // return true

    // LEAST OPTIIMSED
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
