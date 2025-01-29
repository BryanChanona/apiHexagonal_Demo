package application

import "example/apiHexa/src/products/domain"

type ViewAllProducts struct {
	db domain.IProduct
}


func NewViewAllProducts (db domain.IProduct) *ViewAllProducts{
	return &ViewAllProducts{db:db}
}

func (vp *ViewAllProducts) Execute()([]domain.Product, error){
	return vp.db.GetAll()
}
