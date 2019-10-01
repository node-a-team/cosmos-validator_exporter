package config

import (

        t "github.com/node-a-team/cosmos-validator_exporter/types"

	"fmt"
        "time"
	"os"
	"log"

        "github.com/BurntSushi/toml"
)


func Init() {

	var config = readConfig()

        fmt.Printf("\033[7m##### %s #####\033[0m\n", config.Title)

        t.RpcServer = config.Rpc.Address
        t.RestServer = config.Rest_server.Address

	t.OperatorAddr = config.Validator_info.OperatorAddress
	t.ExporterListenPort = config.Option.ExporterListenPort
	t.OutputPrint = config.Option.OutputPrint
        t.Bech32MainPrefix = config.Network

        fmt.Println("\n[ Your Info ]")
        fmt.Println("- Network:", config.Network)
        fmt.Println("- RPC Server Address:", config.Rpc.Address)
        fmt.Println("- Rest Server Address:", config.Rest_server.Address)
        fmt.Println("- Validator Operator Address:", config.Validator_info.OperatorAddress)
        fmt.Println("- Exporter Listen Port:", config.Option.ExporterListenPort)
        fmt.Printf("- Output Print: %v\n\n", config.Option.OutputPrint)


	fmt.Printf("\033[32m## Start Exporter\033[0m\n\n")

	fmt.Printf("\033[34m## This exporter was created by \"Node A-Team\"\n")
	fmt.Printf("## Validator: https://www.mintscan.io/validators/cosmosvaloper14l0fp639yudfl46zauvv8rkzjgd4u0zk2aseys\n")
	fmt.Printf("## Git: https://github.com/node-a-team/cosmos-validator_exporter.git\033[0m\n\n")

        time.Sleep(1*time.Second)
}

func readConfig() t.Config {

	var configfile string = "config.toml"
	var config t.Config

	_, err := os.Stat(configfile)

	if err != nil {
                log.Fatal("Config file is missing: ", configfile)
        }

        if _, err := toml.DecodeFile(configfile, &config); err != nil {
                log.Fatal(err)
        }

        return config

}
