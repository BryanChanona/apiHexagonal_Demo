package application

import "example/apiHexa/src/products/domain"


type ViewProduct struct {
	db domain.IProduct
}


 func  NewViewProductById(db domain.IProduct) *ViewProduct{
	return &ViewProduct{db: db}
 }

 func (viewProduct *ViewProduct) Execute(id int32)(domain.Product, error){
	return viewProduct.db.GetById(id)

 }