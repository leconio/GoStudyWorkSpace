package main

import (
	"log"
	"os"
	// 必须加入下面代码，否则matcher包不会初始化
	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("hello")
}
