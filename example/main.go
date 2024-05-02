package main

import (
	"fmt"
	"log"
	"time"

	"github.com/i14t/go-fileprobe"
)

func main() {
	fp := fileprobe.NewHandler()
	if err := fp.Create(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File probe is exists: %v\n", fp.Exists())

	duration := 5 * time.Second
	for remaining := duration; remaining >= 0; remaining -= time.Second {
		fmt.Printf("Time remaining: %d seconds\n", remaining/time.Second)
		time.Sleep(time.Second)
	}

	if err := fp.Remove(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File probe is exists: %v\n", fp.Exists())
}
