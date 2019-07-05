package models

import "strings"

//import "fmt"

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		//fmt.Println("tag: ", tag)
		tagList := strings.Split(tag, "&")
		//fmt.Println("tagList: ", tagList)
		for _, value := range tagList {
			tagsMap[value]++
		}
	}

	return tagsMap
}
