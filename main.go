package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	restful "github.com/hwangseonu/gin-restful"
	"github.com/hwangseonu/paperless.dev/docs"
	"github.com/hwangseonu/paperless.dev/internal/auth"
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/hwangseonu/paperless.dev/internal/resource"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	engine := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	protector := auth.NewProtector()
	protector.RegisterAny("/api/v1/users/:id")
	protector.Register("/api/v1/resumes", http.MethodPost)
	protector.Register("/api/v1/resumes/:id", http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)

	engine.Use(protector.Middleware())
	engine.Use(common.ErrorHandler)

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

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
