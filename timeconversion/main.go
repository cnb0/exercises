package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	t, err := time.Parse("03:04:05PM", s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())

}
