package types

type ConfigParams struct {
	NetworkName string `toml:"network_name"`
	GrpcUrl     string `toml:"grpc_url"`
	WsUrl       string `toml:"ws_url"`
	Query       string `toml:"query"`
	Subscriber  string `toml:"subscriber"`
}
