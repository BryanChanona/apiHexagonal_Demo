package application

import "example/apiHexa/src/products/domain"



type CreateProduct struct {
	db domain.IProduct
}


func NewCreateProduct(db domain.IProduct) *CreateProduct{
	return &CreateProduct{db: db}

}

func (uc *CreateProduct) Execute( product domain.Product)(err error) {
	
	return uc.db.Save(product)
}



