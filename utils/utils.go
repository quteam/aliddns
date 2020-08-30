package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/quteam/aliddns/logger"
)

type ResData struct {
	Code int
	IP   string
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Env(key string, defaultValue ...string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}
	return v
}

func GetIP() string {
	res, err := httpClient.Get("https://html.quteam.com/ip")
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	defer res.Body.Close()

	target := &ResData{}
	json.NewDecoder(res.Body).Decode(target)

	return target.IP
}
