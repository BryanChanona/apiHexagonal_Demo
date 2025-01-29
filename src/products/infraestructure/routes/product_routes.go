package routes

import (
	"example/apiHexa/src/products/infraestructure/dependencies"
	_ "net/http"

	"github.com/gin-gonic/gin"
)


func ProductRouter( router *gin.Engine){
	routes := router.Group("/products")
	createProductController := dependencies.GetCreateProductController().CreateProduct
	getAllProducts := dependencies.GetCaseViewAllProducts().ViewProduct
	getProductById := dependencies.GetCaseViewProductById().ViewProductById

	
	routes.POST("", createProductController)
	routes.GET("getAllProduct",getAllProducts)
	routes.GET("getProductById/:id",getProductById)
}








