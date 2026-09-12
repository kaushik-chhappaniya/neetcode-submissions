func evalRPN(tokens []string) int {
	id := len(tokens) - 1

	var recursion func() int

	recursion = func() int {
		token := tokens[id]
		id--
		if token != "+" && token != "-" && token != "*" && token != "/" {
			val,  _ := strconv.Atoi(token)
			return val
		}

		right := recursion()
		left := recursion()

		switch token {
			case "+" :
				return left + right
			case "-" :
				return left - right
			case "*" :
				return left * right
			default:
				return left / right
		}
	}
	return recursion()
}
