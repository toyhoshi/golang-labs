package main

import ( 
	"fmt"
	"os"
)

// https://www.cockroachlabs.com/blog/go-file-size/

func main() {
	name := os.Getenv("USERNAME")
	fmt.Printf("Well done %s for having your first Go\n", name)
}
