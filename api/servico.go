package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Teste struct {
  TEST string `json:teste`
}

type Data struct {
  DATA string `json:data`
}

// Random test endpoint
func teste(c *gin.Context) {
  var teste Data
  teste.DATA = obterTeste()
  fmt.Println(teste)
  c.JSON(http.StatusOK, teste)
}

func tiposDocumento(c *gin.Context) {
  response := listarTipdoc()
  c.JSON(http.StatusOK, response)
}

func documentos(c *gin.Context) {
  response := listarHomedocs()
  c.JSON(http.StatusOK, response)
}
