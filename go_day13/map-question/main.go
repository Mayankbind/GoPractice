package main

import "fmt"

func main() {
	mp1 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}

	mp2 := map[string]int{
		"key5": 10,
		"key2": 20,
		"key7": 30,
		"key4": 40,
	}

	mp := mergeMap(mp1, mp2)

	for key, value := range mp {
		fmt.Println(key, ":", value)
	}
}

func removeDuplicate(words []string) []string {

	var resultSlice []string
	seen := make(map[string]bool)

	for _, val := range words {

		if !seen[val] {
			seen[val] = true
			resultSlice = append(resultSlice, val)
		}
	}

	return resultSlice
}

func mergeMap(mp1, mp2 map[string]int) map[string]int {
	resultMap := make(map[string]int)

	for key, value := range mp1 {
		resultMap[key] = value
	}

	for key, value := range mp2 {
		resultMap[key] = value
	}

	return resultMap
}
