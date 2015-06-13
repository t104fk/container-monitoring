package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main()  {
  router := gin.Default()
  v1 := router.Group("/v1")
  {
    v1.GET("/ping", func (c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "pong"})
    })
  }
  router.Run(":3000")
}
