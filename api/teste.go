package api

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

type Teste struct {
  TEST string `json:teste`
}

// Random test endpoint
func teste(c *gin.Context) {
  var teste Teste
  teste.TEST = "TESTE"
  fmt.Println(teste.TEST)
  c.JSON(http.StatusOK, teste)
}
