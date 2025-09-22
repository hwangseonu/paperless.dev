package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/auth"
	"github.com/hwangseonu/paperless.dev/resource"
)

func main() {
	engine := gin.Default()

	protector := auth.NewProtector()
	protector.RegisterAny("/api/v1/users/:id")
	protector.Register("/api/v1/resumes", http.MethodPost)
	protector.Register("/api/v1/resumes/:id", http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)

	engine.Use(protector.Middleware())

	api := restful.NewAPI("/api/v1")
	{
		user := resource.NewUser()
		resume := resource.NewResume()
		api.RegisterResource("/users", user)
		api.RegisterResource("/resumes", resume)
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
