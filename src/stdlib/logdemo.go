package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	log.Println("message")

	log.Fatal("fatal")

	log.Panic("panic")
}
