package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello, world! \\o/\n")

	currentLocalTime := time.Now().Local()
	fmt.Println("The Current time is ", currentLocalTime.Format("2006-01-02 MST 15:04:05"))

	currentUTCtime := time.Now().UTC()
	fmt.Println("The Current time is ", currentUTCtime.Format("2006-01-02 MST 15:04:05"))
}
