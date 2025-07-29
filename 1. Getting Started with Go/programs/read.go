package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the struct
type Name struct {
	fname string
	lname string
}

func main() {
	// Prompt user for file name
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scan(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)

	// Read each line and parse first and last name
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)

		if len(parts) == 2 {
			fname := truncate(parts[0], 20)
			lname := truncate(parts[1], 20)
			person := Name{fname: fname, lname: lname}
			names = append(names, person)
		} else {
			fmt.Println("Skipping malformed line:", line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print all names
	fmt.Println("\nNames found in file:")
	for _, n := range names {
		fmt.Printf("First Name: %-20s Last Name: %-20s\n", n.fname, n.lname)
	}
}

// Helper function to truncate string to 20 characters
func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	return s
}
