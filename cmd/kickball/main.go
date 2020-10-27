package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)
	kickball := func(player string) {
		for {
			fmt.Println(<-ball, "传球")
			time.Sleep(time.Second)
			ball <- player
		}
	}
	go kickball("张三")
	go kickball("李四")
	go kickball("王二麻子")
	go kickball("刘二爷")

	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}
