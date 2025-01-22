package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("Enter your name: ")

    // Read user input
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')

    // Clean up the input and print a greeting
    name = name[:len(name)-1] 
    fmt.Printf("Hello, %s! Welcome to the Go world.\n", name)
}
