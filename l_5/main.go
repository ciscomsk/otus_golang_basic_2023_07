package main

import (
	"fmt"
	"github.com/ciscomsk/otus_golang_basic_l5_util/util"
)

func main() {
	a := 1
	b := 10
	//fmt.Println(a + b) // v1 - local

	// v2 - local package - import "sum/util"
	// v3 - github repo
	c := util.Sum(a, b)
	fmt.Println(c)
}
