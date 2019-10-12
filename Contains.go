package main

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ContainsItem (arr [](*Item), id string) (*Item) {
	for _, a := range arr {
		if (a.Id == id) {
			return a
		}
	}

	return nil
}
