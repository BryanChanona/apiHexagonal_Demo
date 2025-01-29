package main

import (
	"example/apiHexa/src/products/infraestructure/dependencies"
	"example/apiHexa/src/products/infraestructure/routes"
	
	"github.com/gin-gonic/gin"
)

func main() {
	
	dependencies.Init()
	r := gin.Default()
	routes.ProductRouter(r)
	r.Run()

}