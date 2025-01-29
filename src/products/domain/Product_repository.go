package domain

type IProduct interface {
	Save(product Product) (err error)
	GetAll()([]Product,error)
	GetById(id int32)(Product, error)
	//DeleteProduct(id int32)(error)
}