package main

func checkBool(arr []int, start, end int) bool {
	uniqueValues := make(map[int]bool)

	for _, num := range arr {
		uniqueValues[num] = true
	}

	for val := range uniqueValues {
		if val < start || val > end {

			return false
		}
	}

	return true
}
