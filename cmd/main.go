package main

import (
	"log"

	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev"
	"github.com/hwangseonu/paperless.dev/resource"
)

func main() {
	engine := paperless.App()

	user := resource.NewUser()

	api := restful.NewAPI("/api/v1")
	api.RegisterResource("/users", user)
	api.RegisterHandlers(engine)

	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
