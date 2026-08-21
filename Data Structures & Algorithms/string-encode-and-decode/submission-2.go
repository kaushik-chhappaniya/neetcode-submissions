type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var stringsLen []string
	for _, str := range strs {
		stringsLen = append(stringsLen, strconv.Itoa(len(str)))
	}
	return strings.Join(stringsLen, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}
	stringEnc := strings.SplitN(encoded, "#", 2)
	lengths := strings.Split(stringEnc[0], ",")
	var res []string
	i:=0
	for _, sz := range lengths {
		if sz == "" {
			continue
		}
		length, _ := strconv.Atoi(sz)
		res = append(res, stringEnc[1][i:i+length])
		i += length
	}
   return res
	// start, end := 0, 0
	// for _, v := range nums {
	// 	n, _ := strconv.Atoi(strings.TrimSpace(v))
	// 	end = end + n
	// 	result = append(result, stringEnc[start:end])
	// 	start = end
	// }
	// return result
}
