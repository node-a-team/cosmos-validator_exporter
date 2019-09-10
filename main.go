package main

import (
	"log"
	"net/http"

        t "github.com/node-a-team/cosmos-validator_exporter/types"
        config "github.com/node-a-team/cosmos-validator_exporter/function/config"
        exporter "github.com/node-a-team/cosmos-validator_exporter/function/exporter"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (


)



func main() {

        http.Handle("/metrics", promhttp.Handler())

	config.Init()

	go exporter.Exporter()

	 log.Fatal(http.ListenAndServe(":"+t.ExporterListenPort, nil))
 }
