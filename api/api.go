package api

import (
  "github.com/gin-gonic/gin"
)

// This is fucking stupid
// Exported functions has the first letter upper-cased
func Run() {
  router := gin.Default()

  //router.GET("/teste", teste)
  rotas(router)

  router.Run("localhost:8081")
}
