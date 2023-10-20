package main

import (
	"fmt"

	"github.com/cassioglay/go-rest-api/database"
	"github.com/cassioglay/go-rest-api/models"
	"github.com/cassioglay/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
		{Id: 3, Nome: "Nome 3", Historia: "Historia 3"},
	}

	database.ConectaComBancoDeDados()

	fmt.Print("Server running...")
	routes.HandleRequest()
}
