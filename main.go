package main

import (
	"time"

	"github.com/quteam/aliddns/utils"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			println(utils.GetIP())
		}
	}()
	select {}
	// time.Sleep(time.Minute)
}
