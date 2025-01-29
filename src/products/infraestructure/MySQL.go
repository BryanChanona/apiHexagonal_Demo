package infraestructure

import (
	"database/sql"
	"example/apiHexa/src/products/domain"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(product domain.Product) (err error) {
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO product (idproduct, name, price) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	defer sentenciaPreparada.Exec()

	_, err = sentenciaPreparada.Exec(product.Id, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil

}
func (mysql *MySQL) GetAll() ([]domain.Product, error) {
	fmt.Println("Lista de productos")
	data, err := mysql.db.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var products []domain.Product
	// Itera sobre todas las filas devueltas por la consulta
	for data.Next() {
		var product domain.Product
		err := data.Scan(&product.Id,&product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	if err := data.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (mysql *MySQL) GetById(id int32)(domain.Product,error){
	var productById domain.Product
	//Nuestra consulta sql
	 query := "SELECT idproduct, name, price FROM product WHERE idproduct = ?"
	//Ejecutamos la consulta con el id proporcionado, nos devuelve la fila en caso de que la encuentre
	 row := mysql.db.QueryRow(query,id)
	
	 err := row.Scan(&productById.Id,&productById.Name, &productById.Price) //mapea los valores
    if err != nil {
        if err == sql.ErrNoRows {
            return productById, fmt.Errorf("producto con ID %d no encontrado", id)
        }
        return productById, err // Retornar cualquier otro error
    }


    return productById, nil
	 }

 func (mysql *MySQL) deleteProduct (id int32) error{
	

	query := "DELETE FROM product WHERE idproduct = ?"
	//Exec para delete
	result, err := mysql.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Verificar si realmente se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el producto con ID %d", id)
	}
	fmt.Println("Producto eliminado correctamente")
	return nil


 }



	 



