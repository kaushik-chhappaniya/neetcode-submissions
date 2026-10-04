func evalRPN(tokens []string) int {
	var st []int
	for _, v := range tokens {
		switch v {
			case "+","-","/","*":
				b := st[len(st)-1]
				a := st[len(st)-2]
				st = st[:len(st)-2]
				switch v{
					case "+":
						st = append(st, a+b)
					case "-":
						st = append(st, a-b)
					case "*":
						st = append(st, a*b)
					case "/":
						st = append(st, a/b)
				}
			default:
			    num,_ := strconv.Atoi(v)
				st = append(st, num)
		}
	}
	return st[0] 
}
