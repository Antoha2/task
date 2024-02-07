package config

import "google.golang.org/grpc"

//const MQURL = "amqp://guest:guest@localhost:8183/"

type Config struct {
	DB   DBConfig
	Auth AuthConfig
	MQ   MQConfig
	GRPC GRPCConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

type AuthConfig struct {
	Host string
	Port int
}

type MQConfig struct {
	Host string
}

type GRPCConfig struct {
	Conn *grpc.ClientConn
}

var GRPCConn *grpc.ClientConn

func GetConfig() *Config {

	return &Config{
		DB: DBConfig{
			User:     "todoadmin",
			Password: "tododo",
			Host:     "postgres",
			Port:     5432,
			Dbname:   "tododb",
			Sslmode:  "",
		},
		Auth: AuthConfig{

			Host: "auth",
			Port: 8180,
		},
		MQ: MQConfig{

			Host: "amqp://guest:guest@rabbitmq:5672/",
		},
	}

}
