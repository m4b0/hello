package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello, world! *\n")

	current_local_time := time.Now().Local()
	fmt.Println("The Current time is ", current_local_time.Format("2006-01-02"))

	current_utc_time := time.Now().UTC()
	fmt.Println("The Current time is ", current_utc_time.Format("2006-01-02 MST"))
}
