package main

import (
	"fmt"
	"log"
	"main/config"
	"net"
)

func main() {
	config.CONFIG.Init("/home/wyc/code/go/dhtgo/config/config.yaml")
	fmt.Println(config.CONFIG)
	addr, err := net.ResolveUDPAddr("udp", "router.bittorrent.com:6881")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	fmt.Println(addr.IP)
	//listenAddr := new(net.UDPAddr)
	//listener, err := net.ListenUDP("udp", listenAddr)
	//fmt.Println(listener)
}
