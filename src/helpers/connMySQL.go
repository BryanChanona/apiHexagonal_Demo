package helpers

import (
	"database/sql" //Interactuar con base de datos
	"fmt"
	"time"
	"os"
	"github.com/joho/godotenv"
	"log"
	_"github.com/go-sql-driver/mysql"
)
func ConnectDB() (db *sql.DB, err error) {

	if err := godotenv.Load(); err != nil {
        log.Println("No se pudo cargar el archivo .env, verificando variables de entorno...")
    }
	usuario := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	nameDB := os.Getenv("DB_NAME")
	

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", usuario, password, host, nameDB)
	fmt.Println("Cadena de conexión:", dsn)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error al abrir conexión: %s\n", err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		fmt.Printf("Error al verificar la conexión: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
