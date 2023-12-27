package modelos

import (
	"Api1/database"
	"time"
)

// ********** definiendo categoria **********
type Categoria struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
	Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria

// *************** fin ************************

// ********** definiendo producto **********
type Producto struct {
	Id          uint      `json:"id"`
	Nombre      string    `gorm:"type:varchar(100)" json:"nombre"`
	Slug        string    `gorm:"type:varchar(100)" json:"slug"`
	Precio      int       `json:"precio"`
	Stock       int       `json:"stock"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	CategoriaID uint      `json:"categoria_id"` // esta solo crea el campo en la tabla
	Categoria   Categoria `json:"categoria"`    // esto crea la relacion entre tablas
}

type Productos []Producto

// *************** fin ************************

// creando las tablas en la base de datos
func Migraciones() {
	database.Database.AutoMigrate(&Categoria{})
	database.Database.AutoMigrate(&Producto{})
}
