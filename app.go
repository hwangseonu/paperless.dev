package paperless

import (
	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var conf *Config

type Config struct {
	MongoURI string
}

func init() {
	app = gin.Default()
	//conf = &Config{
	//	MongoURI: os.Getenv("MONGO_URI"),
	//}

	conf = &Config{
		MongoURI: "mongodb://localhost:27017",
	}
}

func App() *gin.Engine {
	return app
}

func GetConfig() *Config {
	return conf
}
