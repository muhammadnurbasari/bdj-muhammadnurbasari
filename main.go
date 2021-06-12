package main

import (
	"bdj-muhammadnurbasari/app/registry"
	"bdj-muhammadnurbasari/docs"
)

// !Important : Comments below are formatted as it is to be-
// read by Swagger tools (Swaggo)
// @title SERVICE BDJ
// @version 0.1.5
// @description Service bdj
// @BasePath
// @contact.name Muhammadnurbasari
// @contact.email m.nurbasari@gmail.com
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	appRegistry := registry.NewAppRegistry()
	appRegistry.StartServer()
}
