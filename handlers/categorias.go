package handlers

import (
	"Api1/database"
	"Api1/dto"
	"Api1/modelos"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ListarCategorias(response http.ResponseWriter, request *http.Request) {

	// Configura el encabezado de la respuesta HTTP para indicar que el contenido es JSON
	response.Header().Set("Content-Type", "application/json")

	// Crea una variable llamada 'datos' del tipo 'modelos.Categorias'
	datos := modelos.Categorias{}

	// Realiza una consulta a la base de datos para obtener todas las categorías, ordenadas por ID de forma descendente (últimas primero)
	//database.Database.Find(&datos)
	database.Database.Order("id desc").Find(&datos)

	// Configura el código de estado de la respuesta HTTP como 200 (OK)
	response.WriteHeader(http.StatusOK)

	// Codifica los datos obtenidos de la base de datos en formato JSON y envía la respuesta al cliente
	json.NewEncoder(response).Encode(datos)
}

func ListarCategorias_por_id(response http.ResponseWriter, request *http.Request) {
	// Configura el encabezado de la respuesta HTTP para indicar que el contenido es JSON
	response.Header().Set("Content-Type", "application/json")

	// Obtiene las variables de la solicitud (en este caso, el parámetro "id")
	vars := mux.Vars(request)

	// Declara una variable llamada 'datos' del tipo 'modelos.Categorias'
	datos := modelos.Categoria{}

	// Intenta buscar el primer registro en la base de datos con el ID proporcionado
	if err := database.Database.First(&datos, vars["id"]); err.Error != nil {

		// Si hay un error al buscar el recurso, construye una respuesta de error
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "No se encuentra el recurso!",
		}
		// Configura el código de estado de la respuesta HTTP como 404 (Not Found)
		response.WriteHeader(http.StatusNotFound)

		// Codifica la respuesta de error en formato JSON y envíala al cliente
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		// Si se encuentra el recurso, configura el código de estado como 200 (OK)
		response.WriteHeader(http.StatusOK)

		// Codifica los datos obtenidos de la base de datos en formato JSON y envía la respuesta al cliente
		json.NewEncoder(response).Encode(datos)
	}
}

func CrearCategoria(response http.ResponseWriter, request *http.Request) {

	// Declara una variable llamada 'categoria' del tipo 'dto.CategoriaDTO'
	var categoria dto.CategoriaDTO

	// Intenta decodificar el cuerpo de la solicitud JSON en la variable 'categoria'
	if err := json.NewDecoder(request.Body).Decode(&categoria); err != nil {

		// Si hay un error al decodificar la solicitud, construye una respuesta de error
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrio un error inesperado!",
		}

		// Configura el código de estado de la respuesta HTTP como 400 (Bad Request)
		response.WriteHeader(http.StatusBadRequest)

		// Codifica la respuesta de error en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	// Crea una instancia del modelo de datos 'Categoria' con los datos proporcionados
	datos := modelos.Categoria{Nombre: categoria.Nombre, Slug: categoria.Slug}

	// Guarda la nueva categoría en la base de datos
	database.Database.Save(&datos)

	// Construye una respuesta de éxito
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creo el registro exitosamente!",
	}
	// Configura el encabezado de la respuesta HTTP como JSON y el código de estado como 201 (Created)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	// Codifica la respuesta de éxito en formato JSON y la envía al cliente
	json.NewEncoder(response).Encode(respuesta)

}

func EditarCategoria(response http.ResponseWriter, request *http.Request) {
	// Configura el encabezado de la respuesta HTTP para indicar que el contenido es JSON
	response.Header().Set("Content-Type", "application/json")

	// Obtiene las variables de la solicitud (en este caso, el parámetro "id")
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	// Declara una variable llamada 'categoria' del tipo 'dto.CategoriaDTO'
	var categoria dto.CategoriaDTO

	// Intenta decodificar el cuerpo de la solicitud JSON en la variable 'categoria'
	if err := json.NewDecoder(request.Body).Decode(&categoria); err != nil {
		// Si hay un error al decodificar la solicitud, construye una respuesta de error
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrio un error inesperado!",
		}
		// Configura el código de estado de la respuesta HTTP como 400 (Bad Request)
		response.WriteHeader(http.StatusBadRequest)

		// Codifica la respuesta de error en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
		return
	}

	// Declara una variable llamada 'datos' del tipo 'modelos.Categoria'
	datos := modelos.Categoria{}

	// Intenta buscar el registro en la base de datos con el ID proporcionado
	if err := database.Database.First(&datos, id); err.Error != nil {
		// Si no se encuentra el recurso, construye una respuesta de error
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso NO disponible!",
		}
		// Configura el código de estado de la respuesta HTTP como 404 (Not Found)
		response.WriteHeader(http.StatusNotFound)
		// Codifica la respuesta de error en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		// Actualiza los datos del registro con la información proporcionada en 'categoria'
		datos.Nombre = categoria.Nombre
		datos.Slug = categoria.Slug

		// Guarda los cambios en la base de datos
		database.Database.Save(&datos)

		// Construye una respuesta de éxito
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se modifico el registro exitosamente!",
		}
		// Configura el código de estado de la respuesta HTTP como 201 (Created)
		response.WriteHeader(http.StatusCreated)
		// Codifica la respuesta de éxito en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
	}

}

func EliminarCategoria(response http.ResponseWriter, request *http.Request) {
	// Obtiene las variables de la solicitud (en este caso, el parámetro "id")
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	// Declara una variable llamada 'datos' del tipo 'modelos.Categoria'
	datos := modelos.Categoria{}

	// Intenta buscar el registro en la base de datos con el ID proporcionado
	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		// Configura el encabezado de la respuesta HTTP como JSON y el código de estado como 404 (Not Found)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)

		// Codifica la respuesta de error en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		// Elimina el registro de la base de datos
		database.Database.Delete(&datos)

		// Construye una respuesta de éxito
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se eliminó el registro exitosamente",
		}
		// Configura el encabezado de la respuesta HTTP como JSON y el código de estado como 200 (OK)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)

		// Codifica la respuesta de éxito en formato JSON y la envía al cliente
		json.NewEncoder(response).Encode(respuesta)
	}
}
