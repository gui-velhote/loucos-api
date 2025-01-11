package api

import (
	"github.com/gin-gonic/gin"
)

type callback func()

// This is fucking stupid
// Exported functions has the first letter upper-cased
func Run() {

  router := gin.Default()
  rotas(router)

  router.Run("localhost:8081")
}
