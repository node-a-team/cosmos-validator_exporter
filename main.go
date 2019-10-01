package main

import (
	"log"
	"net/http"

        t "github.com/node-a-team/cosmos-validator_exporter/types"
        config "github.com/node-a-team/cosmos-validator_exporter/function/config"
        websocket "github.com/node-a-team/cosmos-validator_exporter/function/websocket"
        exporter "github.com/node-a-team/cosmos-validator_exporter/function/exporter"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (


)



func main() {

        http.Handle("/metrics", promhttp.Handler())

	config.Init()
	websocket.OpenSocket()

	go exporter.Exporter()

	 log.Fatal(http.ListenAndServe(":"+t.ExporterListenPort, nil))
 }
