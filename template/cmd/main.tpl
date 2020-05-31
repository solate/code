package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	{{- range $key, $value := .ImportList}}
    "{{$value}}"
    {{- end}}

	_ "github.com/solate/util/project/pprof"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 初始化基础组件
	if err := app.InitComponent(); err != nil {
		log.Fatal(err.Error())
	}

	if err := app.InitBusiness(); err != nil {
		log.Fatal(err.Error())
	}
	app.SystemSignal()
}
