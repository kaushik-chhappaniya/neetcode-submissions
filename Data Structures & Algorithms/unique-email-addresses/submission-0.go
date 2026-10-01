func numUniqueEmails(emails []string) int {
    unique := make(map[string]bool)
	for _, v:=range emails {
		parts := strings.Split(v, "@")
		local , domain := parts[0], parts[1]
		local = strings.Split(local, "+")[0]
		local = strings.ReplaceAll(local, ".", "")
		key := local + "@" + domain
		unique[key] = true
	}
	return len(unique)
}