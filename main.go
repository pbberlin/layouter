package main

import (
	"fmt"
	"math/rand"
	"time"
)

var spf func(format string, a ...interface{}) string = fmt.Sprintf
var pf func(format string, a ...interface{}) (int, error) = fmt.Printf

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Main end")
}
