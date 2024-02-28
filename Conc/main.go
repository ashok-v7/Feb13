//

package main

import "fmt"

func sum(nums []int, c chan int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	c <- sum
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(nums[:len(nums)], c)
	//go sum(nums[len(nums)/2:], c)
	x := <-c

	fmt.Println(x)
}
