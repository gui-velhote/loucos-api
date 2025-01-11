package api

import (
  "github.com/gin-gonic/gin"
)

func rotas(router *gin.Engine) {

  /* Declaração de rotas */
  // GETTERS
  router.GET("/teste", teste)
  router.GET("/listar/tipdoc", tiposDocumento)
  router.GET("/listar/homedocs", documentos)

}
