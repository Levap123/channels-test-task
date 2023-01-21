package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := 10
	nums := make([]int, size)

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	fmt.Println("--------NUMBERS-------")
	for i := 0; i < len(nums); i++ {
		nums[i] = random.Intn(10)
		fmt.Println(nums[i])
	}

	workCh := make(chan int)
	doneCh := make(chan int)

	go numSender(workCh, nums)
	go toSquare(workCh, doneCh)

	fmt.Println("-------SQUARES--------")
	for val := range doneCh {
		fmt.Println(val)
	}
	fmt.Println("----------------------")
}

func numSender(ch chan int, nums []int) {
	for _, num := range nums {
		ch <- num
	}
	defer close(ch)
}

func toSquare(workCh, doneCh chan int) {
	for val := range workCh {
		doneCh <- val * val
	}
	defer close(doneCh)
}
