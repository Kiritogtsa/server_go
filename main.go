package main

import "github.com/Kiritogtsa/server_go/routes"

func main() {
	var server routes.Serverfunctions = routes.Newserver(8080, "")
	server.Configroutes()
	server.Runserver()
}
