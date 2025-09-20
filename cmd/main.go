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

	protected := map[string][]string{
		http.MethodGet:    {"/api/v1/users/:id"},
		http.MethodPost:   {"/api/v1/users/:id"},
		http.MethodPut:    {"/api/v1/users/:id"},
		http.MethodPatch:  {"/api/v1/users/:id"},
		http.MethodDelete: {"/api/v1/users/:id"},
	}

	engine.Use(auth.AccessTokenMiddleware(protected))

	api := restful.NewAPI("/api/v1")
	{

		user := resource.NewUser()
		api.RegisterResource("/users", user)
		api.RegisterHandlers(engine)
	}

	authGroup := engine.Group("/api/v1/auth")
	{
		authGroup.POST("/login", auth.LoginHandler)
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
