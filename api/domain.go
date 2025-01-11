package api

import (
	"context"
	"fmt"
	"log"
)

type Tipdoc struct {
  CODTIPDOC int     `json:codtipdoc`
  TIPDOC    string  `json:tipdoc`
}

type Homedocs struct {
  CODDOC int `json:coddoc`
  NOMARQDOC string `json:nomarqdoc`
  CODTIPDOC int `json:codtipdoc`
}

func obterTeste() string {
  db := databaseConnection()
  var teste string
  err := db.QueryRow(context.Background(), "select 'Teste Ok!'").Scan(&teste)
  if err != nil {
    log.Fatal(err.Error())
  }
  db.Close()
  return teste
}

func listarTipdoc() []Tipdoc {
  db := databaseConnection()
  var response []Tipdoc
  rows, err := db.Query(context.Background(), "select * from tipdoc")

  if err != nil {
    log.Fatal("Erro ao obter tipdoc")
    db.Close()
  }

  for rows.Next() {
    var tempResp Tipdoc
    if err := rows.Scan(&tempResp.CODTIPDOC, &tempResp.TIPDOC); err != nil {
      log.Fatal("Erro ao scanear resultados!")
      db.Close()
      return nil
    }
    response = append(response, tempResp)
  }

  db.Close()
  fmt.Println(response)

  return response
}

func listarHomedocs() []Homedocs {
  db := databaseConnection()

  var response []Homedocs
  rows, err := db.Query(context.Background(), "SELECT * FROM HOMEDOCS")
  if err != nil {
    log.Fatal("Não foi possível listar os documentos")
    db.Close()
    return nil
  }

  for rows.Next() {
    var tempResp Homedocs
    if err := rows.Scan(&tempResp.CODTIPDOC, &tempResp.CODDOC, &tempResp.NOMARQDOC); err != nil {
      log.Fatal("Erro ao obter linha")
      db.Close()
      return nil
    }
    response = append(response, tempResp)
  }

  db.Close()
  fmt.Println(response)
  return response
}
