package config

func GetPort() string {
	return port
}

// change to env vars maybe
var port string = "8080"
var host string = "localhost"
