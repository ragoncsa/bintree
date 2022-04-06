package main

import (
	"bintree/bintree"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := []int{}
	for n := 0; n < 20; n++ {
		data = append(data, int(r.Float64()*100))
	}
	fmt.Printf("INPUT: %#v\n", data)
	t := bintree.BintreeFromSlice(data)
	fmt.Println(t.Render())
}
