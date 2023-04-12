package main

import (
	"fmt"
	"go-algirithms/application/concurrent"
	"go-algirithms/application/sort"
	"math/rand"
	"time"
)

func main() {
	callReadFileCountWord()
}

func callReadFileCountWord() {
	concurrent.ReadFileCountWord()
}

func callSelectSort() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 20)
	for i := 0; i < cap(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Printf("nums: %d\n", sort.SelectSort(nums))
}
