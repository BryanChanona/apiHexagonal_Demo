package controller

import (
	"example/apiHexa/src/products/application"
	"example/apiHexa/src/products/domain"
		_"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

)
//Estrutura CreateProductController -->Esta estructura representa un controlador para manejar las solicitudes relacionadas con la creación de productos.
type CreateProductController struct{
	createProduct *application.CreateProduct
}

//Constructor --> permite inyectar la dependencia del caso de uso (application.CreateProduct) al controlador. 	
func NewCreateProductController(uc *application.CreateProduct) *CreateProductController{
	return &CreateProductController{createProduct: uc}
}

func ( c *CreateProductController)CreateProduct(ctx *gin.Context){
	var product domain.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.createProduct.Execute(product)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Product created"})
	}



}



