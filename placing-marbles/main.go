package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scanf("%f", &a)
	b := math.Floor(a / 100)
	c := math.Floor((a - b*100) / 10)
	d := a - b*100 - c*10
	fmt.Println(b + c + d)
}
