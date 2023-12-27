package main

import (
	"Api1/handlers"
	"Api1/modelos"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// migrar la bd
	modelos.Migraciones() // ayuda a crear las tablas en la BD, debemos descomentarla para usarla

	mux := mux.NewRouter()

	prefijo := "/api/v1/"
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo_get_json", handlers.Ejemplo_get_json).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo_get_query_string", handlers.Ejemplo_get_query_string).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo_parametros/{id:[0-9]+}", handlers.Ejemplo_get_con_parametros).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo_post", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo_put/{id:[0-9]+}/{nombre}", handlers.Ejemplo_put).Methods("PUT") // recibe un numero y un string en la url
	mux.HandleFunc(prefijo+"ejemplo_delete/{id:[0-9]+}", handlers.Ejemplo_delete).Methods("DELETE")
	mux.HandleFunc(prefijo+"upload", handlers.Ejemplo_upload).Methods("POST")          //subida de archivos
	mux.HandleFunc(prefijo+"ejemplo_ver_foto", handlers.EjemploVerFoto).Methods("GET") //subida de archivos

	// categorias
	mux.HandleFunc(prefijo+"categorias", handlers.ListarCategorias).Methods("GET")                         //mostrar las categorias
	mux.HandleFunc(prefijo+"categoriasById/{id:[0-9]+}", handlers.ListarCategorias_por_id).Methods("GET")  //mostrar las categorias_por_id
	mux.HandleFunc(prefijo+"crear_categoria", handlers.CrearCategoria).Methods("POST")                     //crear categorias
	mux.HandleFunc(prefijo+"editar_categoria/{id:[0-9]+}", handlers.EditarCategoria).Methods("PUT")        //edita categoria
	mux.HandleFunc(prefijo+"eliminar_categoria/{id:[0-9]+}", handlers.EliminarCategoria).Methods("DELETE") //edita categoria

	mux.HandleFunc(prefijo+"productos", handlers.ListarProductos).Methods("GET")
	mux.HandleFunc(prefijo+"productosById/{id:[0-9]+}", handlers.ListarProductos_por_id).Methods("GET")  //mostrar las categorias_por_id
	mux.HandleFunc(prefijo+"crear_producto", handlers.CrearProducto).Methods("POST")                     //crear categorias
	mux.HandleFunc(prefijo+"editar_producto/{id:[0-9]+}", handlers.EditarProducto).Methods("PUT")        //edita categoria
	mux.HandleFunc(prefijo+"eliminar_producto/{id:[0-9]+}", handlers.EliminarProducto).Methods("DELETE") //edita categoria

	//cors
	handler := cors.AllowAll().Handler(mux)
	//log.Fatal(http.ListenAndServe(":8080", mux))
	log.Fatal(http.ListenAndServe(":8080", handler))
}
