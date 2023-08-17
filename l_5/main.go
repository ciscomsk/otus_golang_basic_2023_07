package main

import (
	"fmt"
	"sum/util"
)

func main() {
	a := 1
	b := 10
	//fmt.Println(a + b)

	c := util.Sum(a, b)
	fmt.Println(c)
}
