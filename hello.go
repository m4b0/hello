package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello, world! *\n")

	current_time := time.Now().Local()
	fmt.Println("The Current time is ", current_time.Format("2006-01-02"))
	
	current_time := time.Now().UTC()
    fmt.Println("The Current time is ", current_time.Format("2006-01-02 MST"))
}
