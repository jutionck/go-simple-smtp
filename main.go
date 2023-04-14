package main

import "github.com/jutionck/go-simple-smtp/delivery"

func main() {
	delivery.NewServer().Run()
}
