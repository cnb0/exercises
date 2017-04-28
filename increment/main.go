package main

import (
	"fmt"
	"strconv"
)

var tags []string

func main() {
	var newTags []string
	tags = []string{"1", "latest", "foo"}
	for _, t := range tags {
		i, err := strconv.Atoi(t)
		if err != nil {
			newTags = append(newTags, t)
			continue
		}
		i++
		newTags = append(newTags, strconv.Itoa(i))
	}
	fmt.Println(newTags)
}
