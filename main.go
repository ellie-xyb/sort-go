package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func getData() []int {
	data := []int{}
	for i := 0; i < 10000000; i++ {
		data = append(data, rand.Intn(4294998))
	}
	return data
}

func main() {
	// fmt.Println("hello word")
	data := getData()
	sort.Ints(data)
	fmt.Println(data[0])
}
