package common

import "os"

var conf *Config

type Config struct {
	MongoURI  string
	JwtSecret string
}

func init() {
	conf = &Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func GetConfig() *Config {
	return conf
}
