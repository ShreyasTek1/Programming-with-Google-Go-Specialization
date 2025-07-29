// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {

// 	nums:=make([]int, 0, 3)

// 	for {
// 		fmt.Print("Enter an integer (or 'X' to quit): ")
// 		var input string
// 		fmt.Scanln(&input)

// 		if input == "X" || input == "x" {
// 			fmt.Println("Exiting program.")
// 			break
// 		}

// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a number or 'X'.")
// 			continue
// 		}

// 		nums = append(nums, num)
// 		sort.Ints(nums)
// 		fmt.Println("Sorted slice:", nums)
// 	}
// }
