package prometheus

import (

//        "fmt"
        "os/exec"
//        "encoding/json"
//        "regexp"
)

var (
	addr string = "localhost:26660"
	PMetric string
)


func GetPrometheusMetric() {

//        check := "tendermint_consensus_block_interval_seconds"
//        check = "tendermint_p2p_peer_send_bytes_total"

        cmd := "curl -s -XGET " +addr
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	PMetric = string(out)

/*
        re, _ := regexp.Compile(check +"{[0-9a-zA-Z-_.,\\\"=]+}" +"\\s" + "[0-9]+(.)[0-9]+")
        str := re.FindString(string(out))
        fmt.Println(str)
        fmt.Println(string(out))
*/
}

