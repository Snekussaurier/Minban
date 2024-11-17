package main

import (
	"github.com/snekussaurier/minban-backend/controller"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/routes"
)

func main() {
	database.InitializeDB()
	controller.CreateDefaultUser()
	routes.SetupRouter().Run(":9916")
}
