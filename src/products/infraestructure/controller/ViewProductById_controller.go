package controller

import (
	"example/apiHexa/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewProductByIdController struct {
	viewProductById *application.ViewProduct
}

func NewViewByIdController(uc *application.ViewProduct) *ViewProductByIdController{
	return &ViewProductByIdController{viewProductById: uc}

}

func (controllerViewProductById *ViewProductByIdController) ViewProductById(ctx *gin.Context){
	//Recibimos el id de la url
	idParam := ctx.Param("id")
	//Convertimos a entero
	id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
	product, err := controllerViewProductById.viewProductById.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product": product})
}