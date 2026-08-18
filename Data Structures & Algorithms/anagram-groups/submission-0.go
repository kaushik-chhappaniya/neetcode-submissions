func groupAnagrams(strs []string) [][]string {
    hm := make(map[string][]string)

    for _,v := range(strs) {
        sortedStr := sorted(v)
        hm[sortedStr] = append(hm[sortedStr], v)
    }
    var result [][]string
    for _, stringGroup := range hm {
        result = append(result, stringGroup)
    }
        return result


}

func sorted(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i,j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}