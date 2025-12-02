package main

import "fmt"

func main(){
	var sampleMap = map[string]interface{}{
		"key1":1,
		"key2":[]map[string]interface{}{}}
	fmt.Println(sampleMap["key3"])
}
