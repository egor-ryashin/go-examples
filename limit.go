package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func limit() {
	limiter := rate.NewLimiter(rate.Every(10*time.Second), 10)
	for i := 0; i < 20; i++ {
		fmt.Printf("%t\n", limiter.Allow())
	}
}
