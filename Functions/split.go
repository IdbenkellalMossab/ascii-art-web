package function

func Split(str string) []string {
	slice := []string{}
	newStr := ""

	for i := 0; i < len(str); i++ {
		// check that \ followed by an n.
		if i < len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			// if yes we check if there is something before \n
			if newStr != "" {
				slice = append(slice, newStr)
				newStr = ""
			}
			// instead of /n we add a "" to avoid the problem of duplicating \n
			slice = append(slice, "")
			// here we skip the next element which is n
			i += 1

		} else { // otherwise we add directly to the string
			newStr += string(str[i])
		}
	}

	if newStr != "" {
		slice = append(slice, newStr)
	}
	return slice
}