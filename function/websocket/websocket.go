package websocket

import (
   t "github.com/node-a-team/cosmos-validator_exporter/types"
   tmclient "github.com/tendermint/tendermint/rpc/client"
//   "encoding/json"
   "fmt"
//   "os/exec"
//   "strconv"

)


func OpenSocket() {
        t.Client = tmclient.NewHTTP("tcp://"+t.RpcServer, "/websocket")

        err := t.Client.Start()
        if err != nil {
                // handle error
		fmt.Println("######## RPC/websockets Connection Error ########")
        }
        defer t.Client.Stop()

        fmt.Println("####### RPC/websockets Connect -> tcp://"+t.RpcServer +" ########")
}



