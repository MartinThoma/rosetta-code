package main

import "log"

func main() {
	var a = 1
	var b = 3
	var c = a / b
	log.Println(c) // gives 0 - no warning!
	var d = float64(a) / float64(b)
	log.Println(d) // gives 0.333333333
}
