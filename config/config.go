package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/quteam/aliddns/logger"
	"github.com/quteam/aliddns/utils"
)

type Config struct {
	Region          string `json:"region"`
	AccessKey       string `json:"accessKey"`
	AccessKeySecret string `json:"accessKeySecret"`
	Domain          string `json:"domain"`
	RR              string `json:"rr"`
	Interval        int64  `json:"interval"`
}

var Base = &Config{
	Region:   "cn-hangzhou",
	Interval: 20,
}

func loadConfig(filename string, config *Config) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	loadConfig("./config.json", Base)
	if accessKey := utils.Env("ACCESS_KEY"); accessKey != "" {
		Base.AccessKey = accessKey
	}
	if accessKeySecret := utils.Env("ACCESS_KEY_SECRET"); accessKeySecret != "" {
		Base.AccessKeySecret = accessKeySecret
	}
	if domain := utils.Env("DOMAIN"); domain != "" {
		Base.Domain = domain
	}
	if rr := utils.Env("RR"); rr != "" {
		Base.RR = rr
	}
	if interval := utils.Env("INTERVAL"); interval != "" {
		val, err := strconv.ParseInt(interval, 10, 64)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		Base.Interval = val
	}
}
