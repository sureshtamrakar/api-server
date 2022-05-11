package main

import (
	"github.com/sureshtamrakar/api-server/routes"
)

func main() {
	r := routes.AddRoutes()
	r.Run()

}
