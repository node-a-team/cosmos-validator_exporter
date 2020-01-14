package config

import (
	"log"

	"github.com/BurntSushi/toml"

	rpc "github.com/node-a-team/cosmos-validator_exporter/getData/rpc"
	rest "github.com/node-a-team/cosmos-validator_exporter/getData/rest"
)

const (
)

var (
	ConfigPath string
	Config	configType
)


type configType struct {

	Title	string	`json:"title"`

	Servers struct {
                Addr struct {
                        RPC	string `json:"rpc"`
                        REST	string `json:"rest"`
                }
        }

	Validator struct {
		OperatorAddr	string	`json:"operatorAddr"`
	}

	Options	struct {
		ListenPort	string	`json:"listenPort"`
	}
}


func Init() string {

	Config = readConfig()

	rpc.Addr = Config.Servers.Addr.RPC
	rest.Addr = Config.Servers.Addr.REST

	rest.OperAddr = Config.Validator.OperatorAddr

	return Config.Options.ListenPort
}

func readConfig() configType {

        var config configType

        if _, err := toml.DecodeFile(ConfigPath +"/config.toml", &config); err != nil{

                log.Fatal("Config file is missing: ", config)
        }

	return config

}
