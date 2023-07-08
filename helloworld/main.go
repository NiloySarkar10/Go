// Creating an executable package (one that creates an exe when "go build" is run).
// Other type is reusable package.
// All main (executable packages) need to have a "main" function.
package main

// Part of Go standard library
import "fmt"

// Function created and declared
func main() {
	fmt.Println("Hello World!")
}
