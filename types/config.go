package types

var (
	RestServer string = "localhost:1317"
	RpcServer string = "localhost:26657"

	OperatorAddr string

	ExporterListenPort string = "29990"
	OutputPrint bool = true
)


type Config struct {

	Title				string	`json:"title"`
	Network				string	`json:"network"`

	Rpc struct {
		Address			string	`json:"address"`
	}
	Rest_server struct {
		Address			string	`json:"address"`
	}
	Validator_info struct {
		OperatorAddress		string	`json:"operatorAddress"`
	}
	Option struct {
		ExporterListenPort	string	`json:"exporterListenPort"`
		OutputPrint		bool	`json:"outputPrint"`
	}

}
