package main

import (
	"fmt"
	"main/config"
	"main/model"
)

func main() {
	config.CONFIG.Init("/home/wyc/code/go/dhtgo/config/config.yaml")
	fmt.Println(config.CONFIG)
	model.Upsert()
}
