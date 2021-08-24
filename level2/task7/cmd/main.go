package main

import "task7/internal/server"

func main() {
	serv := server.NewServer()
	serv.Router.Logger.Fatal(serv.Router.Start(":1323"))
}

