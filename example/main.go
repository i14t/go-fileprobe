package main

import (
	"fmt"
	"time"

	"github.com/i14t/go-fileprobe"
)

func main() {
	fp := fileprobe.NewHandler()
	fp.Create()
	fmt.Printf("File probe is exists: %v\n", fp.Exists())

	duration := 5 * time.Second
	for remaining := duration; remaining >= 0; remaining -= time.Second {
		fmt.Printf("Time remaining: %d seconds\n", remaining/time.Second)
		time.Sleep(time.Second)
	}

	fp.Remove()
	fmt.Printf("File probe is exists: %v\n", fp.Exists())
}
