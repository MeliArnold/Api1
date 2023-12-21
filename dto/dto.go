package dto

type CategoriaDTO struct {
	Nombre string `json:"nombre"`
	Slug   string `json:"slug"`
}

type ProductoDTO struct {
	Nombre      string `json:"nombre"`
	Slug        string `json:"slug"`
	Precio      int    `json:"precio"`
	Stock       int    `json:"stock"`
	Descripcion string `json:"descripcion"`
	CategoriaID uint   `json:"categoria_id"`
}
