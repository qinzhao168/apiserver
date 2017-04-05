package componentconfig

type ApiserverConfig struct {
	HttpAddr string
	HttpPort int
	RpcAddr  string
	RpcPort  int
}

type DockerBuildConfig struct {
	HttpAddr string
	HttpPort int
	RpcAddr  string
	RpcPort  int
}
