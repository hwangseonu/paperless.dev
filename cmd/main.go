package main

import (
	"log"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/resource"
)

func main() {
	engine := gin.Default()

	protected := auth.NewProtected()
	protected.RegisterAny("/api/v1/users/:id")

	engine.Use(protected.Middleware())

	api := restful.NewAPI("/api/v1")
	{
		user := resource.NewUser()
		resume := resource.NewResume()
		api.RegisterResource("/users", user)
		api.RegisterResource("/resume", resume)
		api.RegisterHandlers(&engine.RouterGroup)
	}

	authGroup := engine.Group("/api/v1/auth")
	{
		authGroup.POST("/login", auth.LoginHandler)
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
