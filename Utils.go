package main

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func RemoveIndex(s [](*Item), index int) [](*Item) {
	return append(s[:index], s[(index+1):]...)
}
