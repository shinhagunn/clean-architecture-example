package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Bearer sadfsdfsdf"
	s = strings.TrimPrefix(s, "Bearer ")
	fmt.Println(s)
}
