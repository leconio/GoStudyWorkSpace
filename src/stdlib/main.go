package main

import "os"

func main() {
	//jsondemo.Run3()

	_, err := os.Create("asd.txt")
	if err != nil {
		panic(err)
	}
}
