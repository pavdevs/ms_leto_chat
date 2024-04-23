package server

type Config struct {
	Host string
	Port string
}

func NewConfig(host string, port string) Config {
	return Config{
		Host: host,
		Port: port,
	}
}
