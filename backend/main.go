package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	//init database
	model.InitDb()
	routes.InitRouter()
}
