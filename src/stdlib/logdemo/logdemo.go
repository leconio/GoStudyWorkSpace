package logdemo

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func Run() {
	log.Println("message")

	log.Fatal("fatal")

	log.Panic("panic")
}
