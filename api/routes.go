package api

import (
  "github.com/gin-gonic/gin"
)

func rotas(router *gin.Engine) {

  router.GET("/teste", teste)

}
