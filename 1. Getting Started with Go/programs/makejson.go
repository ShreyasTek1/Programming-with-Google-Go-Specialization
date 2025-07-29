// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter your name: ")
// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)

// 	fmt.Print("Enter your address: ")
// 	address, _ := reader.ReadString('\n')
// 	address = strings.TrimSpace(address)

// 	info := map[string]string{
// 		"name":    name,
// 		"address": address,
// 	}

// 	jsonData, err := json.Marshal(info)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
