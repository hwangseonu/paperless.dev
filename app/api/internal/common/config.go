package common

import (
	"log"
	"os"
	"strconv"
)

var conf *Config

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	MongoURI  string
	JwtSecret string
	RedisAddr string
	SMTP      SMTPConfig
}

func init() {
	mongoURI := os.Getenv("MONGO_URI")
	jwtSecret := os.Getenv("JWT_SECRET")
	redisAddr := os.Getenv("REDIS_ADDR")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	smtpPortNumber, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal(err)
	}

	conf = &Config{
		MongoURI:  mongoURI,
		JwtSecret: jwtSecret,
		RedisAddr: redisAddr,
		SMTP: SMTPConfig{
			Host:     smtpHost,
			Port:     smtpPortNumber,
			Username: smtpUsername,
			Password: smtpPassword,
		},
	}
}

func GetConfig() *Config {
	return conf
}
