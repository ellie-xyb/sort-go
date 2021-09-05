package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getData() []uint32 {
	data := []uint32{}
	for i := 0; i < 10000000; i++ {
		data = append(data, rand.Uint32())
	}
	return data
}

func radixSort(data *[]uint32) {
	out := []uint32{}
	tmp := [][]uint32{}
	for i := 0; i < 256; i++ {
		tmp = append(tmp, []uint32{})
	}

	for byte := 0; byte < 4; byte++ {
		for _, value := range *data {
			b := (value >> (byte * 8)) & 0xFF

			tmp[b] = append(tmp[b], value)
		}

		for b, bucket := range tmp {
			out = append(out, bucket...)
			tmp[b] = []uint32{}
		}

		*data = out
		out = []uint32{}
	}
}

func main() {
	// fmt.Println("hello word")
	data := getData()

	start := time.Now()
	radixSort(&data)
	fmt.Println(time.Since(start))
	fmt.Println(data[0:10])
}
