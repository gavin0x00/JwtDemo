package main

import (
	"github.com/gin-gonic/gin"

	"JwtDemo/api"
	"JwtDemo/middleware/jwt"
)

func main() {
	r := gin.Default()
	r.POST("/login", api.Login)
	r.POST("/register", api.RegisterUser)

	taR := r.Group("/data")
	taR.Use(jwt.JWTAuth())

	{
		taR.GET("/dataByTime", api.GetDataByTime)
	}
	r.Run(":8080")
}
