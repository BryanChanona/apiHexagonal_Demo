package controller

import (
	"example/apiHexa/src/products/application"

	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewAllProductsController struct {
	viewAllProduct *application.ViewAllProducts
}

func NewViewProductsController (uc *application.ViewAllProducts) *ViewAllProductsController{
	return &ViewAllProductsController{viewAllProduct: uc}

}
func (c *ViewAllProductsController)ViewProduct(ctx *gin.Context){
	products, err := c.viewAllProduct.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})

}