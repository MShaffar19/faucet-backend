package config

import (
	"github.com/go-ini/ini"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	PrivateKey      string   `json:"PRIVATEKEY"`
	PublicKey       string   `json:"PUBLICKEY"`
	AccountAddress  string   `json:"ACCOUNTADDRESS"`
	Node            string   `json:"NODE"`
	LCDNode         string   `json:"LCDNODE"`
	Amount          string   `json:"AMOUNT"`
	Origins         []string `json:"ORIGINS"`
	RedisEndpoint   string   `json:"REDISENDPOINT"`
	RedisPassword   string   `json:"REDISPASSWORD"`
	RecaptchaSecret string   `json:"RECAPTCHASECRET"`
	AWSRegion       string   `json:"AWSREGION"`
	Timeout         int64    `json:"TIMEOUT"`
}

func GetConfigFromFile(configFile string) (*Config, error) {
	inicfg, err := ini.Load(configFile)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		PrivateKey:      inicfg.Section("").Key("PRIVATEKEY").String(),
		PublicKey:       inicfg.Section("").Key("PUBLICKEY").String(),
		AccountAddress:  inicfg.Section("").Key("ACCOUNTADDRESS").String(),
		Node:            inicfg.Section("").Key("NODE").String(),
		LCDNode:         inicfg.Section("").Key("LCDNODE").String(),
		Amount:          inicfg.Section("").Key("AMOUNT").String(),
		Origins:         inicfg.Section("").Key("ORIGINS").Strings(","),
		RedisEndpoint:   inicfg.Section("").Key("REDISENDPOINT").String(),
		RedisPassword:   inicfg.Section("").Key("REDISPASSWORD").String(),
		RecaptchaSecret: inicfg.Section("").Key("RECAPTCHASECRET").String(),
		AWSRegion:       inicfg.Section("").Key("AWSREGION").String(),
	}
	cfg.Timeout, err = inicfg.Section("").Key("TIMEOUT").Int64()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetConfigFromENV() (*Config, error) {
	config := Config{
		PrivateKey:      os.Getenv("PRIVATEKEY"),
		PublicKey:       os.Getenv("PUBLICKEY"),
		AccountAddress:  os.Getenv("ACCOUNTADDRESS"),
		Node:            os.Getenv("NODE"),
		LCDNode:         os.Getenv("LCDNODE"),
		Amount:          os.Getenv("AMOUNT"),
		RedisEndpoint:   os.Getenv("REDISENDPOINT"),
		RedisPassword:   os.Getenv("REDISPASSWORD"),
		RecaptchaSecret: os.Getenv("RECAPTCHASECRET"),
		AWSRegion:       os.Getenv("AWSREGION"),
	}

	timeoutString := os.Getenv("TIMEOUT")
	if timeoutString == "" {
		timeoutString = "90"
	}
	timeout, err := strconv.ParseInt(timeoutString, 10, 64)
	if err != nil {
		return nil, err
	}
	config.Timeout = timeout
	// parse comma-separated list of origins
	config.Origins = strings.Split(os.Getenv("ORIGINS"), ",")
	return &config, nil
}
