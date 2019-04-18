package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	RunMode string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("Fail to parse app.ini: ", err)
	}
	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server' : ", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HttpPort = sec.Key("HTTP_PORT").MustString("8080")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
