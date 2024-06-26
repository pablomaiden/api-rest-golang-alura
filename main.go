package main

import (
	"fmt"

	"github.com/pablomaiden/api-rest-golang-alura/database"
	"github.com/pablomaiden/api-rest-golang-alura/models"
	"github.com/pablomaiden/api-rest-golang-alura/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
