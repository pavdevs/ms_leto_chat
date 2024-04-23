package database

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func NewConfig(host string, port string, user string, password string, dbName string) Config {

	return Config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,
	}
}
