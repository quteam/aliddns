package main

import (
	"time"

	"github.com/quteam/aliddns/config"
	"github.com/quteam/aliddns/dns"
	"github.com/quteam/aliddns/logger"
	"github.com/quteam/aliddns/utils"
)

func updateDNS() {
	dnsHandler := dns.New(config.Base.Domain, utils.GetIP(), config.Base.RR)
	if dnsHandler == nil {
		return
	}

	err := dnsHandler.Bind()
	if err != nil {
		logger.Error(err.Error())
	}
}

func main() {
	updateDNS()
	ticker := time.NewTicker(time.Second * time.Duration(config.Base.Interval))
	go func() {
		for _ = range ticker.C {
			updateDNS()
		}
	}()
	select {}
}
