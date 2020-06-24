package controllers

import (
  "github.com/gin-gonic/gin"
  "go-gin-rest-api/models"
  "net/http"
)

func GetAllMovies(c *gin.Context) {
  var movies []models.Movie
  models.DB.Find(&movies)

  c.JSON(http.StatusOK, gin.H{"data": movies})
}

type CreateMovieInput struct {
  Title     string  `json:"title" binding:"required"`
  Director  string  `json:"director" binding:"required"`
}

func CreateMovie(c *gin.Context) {
  var input CreateMovieInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  movie := models.Movie{Title: input.Title, Director: input.Director}
  models.DB.Create(&movie)

  c.JSON(http.StatusOK, gin.H{"data": movie})
}

func GetMovie(c *gin.Context) {
  var movie models.Movie

  if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": movie})
}

type UpdateMovieInput struct {
  Title     string  `json:"title"`
  Director  string  `json:"director"`
}

func UpdateMovie(c *gin.Context) {
  var movie models.Movie

  if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Movie wasn't found!"})
    return
  }

  var input UpdateMovieInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&movie).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": movie})
}

func DeleteMovie(c *gin.Context) {
  var movie models.Movie

  if err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Movie wasn't found!"})
    return
  }

  models.DB.Delete(&movie)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
