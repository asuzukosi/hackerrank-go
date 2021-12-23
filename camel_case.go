package main

//camelCase <- this is a sample of cammel case, we are to count the number of words in the string

func isUppercase(s rune) bool {
	return int(s) >= 'A' && int(s) <= 'Z'
}

func getCamelCaseWordCount(s string) int {
	list := []string{}
	var temp string = ""

	for _, c := range s {
		if isUppercase(c) {
			list = append(list, temp)
			temp = string(c)
		} else {
			temp = temp + string(c)
		}
	}
	list = append(list, temp)
	return len(list)
}
