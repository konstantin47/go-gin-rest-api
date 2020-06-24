package main

import (
  "github.com/gin-gonic/gin"
  "go-gin-rest-api/models"
  "go-gin-rest-api/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/movies", controllers.GetAllMovies)
  r.POST("/movies", controllers.CreateMovie)
  r.GET("/movies/:id", controllers.GetMovie)
  r.PATCH("/movies/:id", controllers.UpdateMovie)
  r.DELETE("/movies/:id", controllers.DeleteMovie)

  r.Run()
}
