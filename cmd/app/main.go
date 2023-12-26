package main

import (
	"com.csdl/demo/config"
	"fmt"
)

func main() {
	var dbConfig = config.OrientDBConfig{}
	var connect = dbConfig.GetConnect()
	fmt.Println("Hello World", connect.GetCurDB())
}
