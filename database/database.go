package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Declara una variable llamada Database y le asigna una función anónima que devuelve un puntero a gorm.DB.
var Database = func() (db *gorm.DB) {
	// Intenta cargar las variables de entorno desde un archivo .env.
	errorVariables := godotenv.Load()
	// Si hay un error durante la carga, lanza una excepción y detiene la ejecución del programa.
	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	// Construye la cadena de conexión a la base de datos MySQL utilizando las variables de entorno.
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Intenta abrir una conexión a la base de datos utilizando GORM y la cadena de conexión construida.
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		// Si hay un error durante la conexión, imprime un mensaje de error y lanza una excepción.
		fmt.Println("Error de conexión")
		panic(err)
	} else {
		// Si la conexión es exitosa, imprime un mensaje indicando la conexión exitosa y devuelve el puntero a gorm.DB.
		fmt.Println("Conexión exitosa")
		return db
	}
}()

// La función anónima se ejecuta inmediatamente después de su definición, estableciendo la conexión a la base de datos
// y asignando el resultado a la variable Database.
