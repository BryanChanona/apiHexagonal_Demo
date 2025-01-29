package dependencies

import (
	"example/apiHexa/src/helpers"
	"example/apiHexa/src/products/application"
	"example/apiHexa/src/products/infraestructure"
	"example/apiHexa/src/products/infraestructure/controller"
	"fmt"

)

var (
	mySQL infraestructure.MySQL
)



func Init() {
	db, err := helpers.ConnectDB()

	if err != nil {
		fmt.Println("Server error")
		return
	}

	defer db.Close()
	mySQL =*infraestructure.NewMySQL(db)
}

//Caso de uso para crear un producto
func GetCreateProductController () *controller.CreateProductController{
	caseCreateProduct := application.NewCreateProduct(&mySQL)
	return controller.NewCreateProductController(caseCreateProduct)

}
//Caso de uso para ver todos los productos

func GetCaseViewAllProducts () *controller.ViewAllProductsController {
	caseViewAllProducts := application.NewViewAllProducts(&mySQL)
	return controller.NewViewProductsController(caseViewAllProducts)
}
 // Caso de uso para ver solo un producto
func GetCaseViewProductById () *controller.ViewProductByIdController {
	caseViewProductbyId := application.NewViewProductById(&mySQL)
	return controller.NewViewByIdController(caseViewProductbyId)
}


