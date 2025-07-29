// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter a String: ")
// 	mname, _ := reader.ReadString('\n')           // Reads full line including spaces
// 	mname = strings.TrimSpace(mname)              // Removes newline and extra spaces or \n would be considere
// 	name := strings.ToLower(mname)                // Convert to lowercase

// 	if len(name) > 0 && strings.HasPrefix(name, "i") && strings.HasSuffix(name, "n") && strings.Contains(name, "a") {
// 		fmt.Println("Found")
// 	} else {
// 		fmt.Println("Not Found")
// 	}
// }
