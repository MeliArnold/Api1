package handlers

import (
	"Api1/dto"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola ejemplo")
}

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

// retorna un json
func Ejemplo_get_json(response http.ResponseWriter, request *http.Request) {
	// Obtener los parámetros de la URL usando mux.Vars
	vars := mux.Vars(request)

	// Establecer el encabezado "Content-Type" en la respuesta como "application/json"
	response.Header().Set("Content-Type", "application/json")

	// Agregar un encabezado personalizado llamado "Tamila" a la respuesta
	response.Header().Add("Tamila", "www.tamila.co")

	// Crear una estructura ResponseGenerico con el estado "ok" y un mensaje que incluye el ID de la solicitud GET
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Get Json | id=" + vars["id"]})

	// Escribir la respuesta JSON en el cuerpo de la respuesta
	fmt.Fprintf(response, string(output))
}

// con query string recuperamos el dato que envian por la url con su nombre
func Ejemplo_get_query_string(response http.ResponseWriter, request *http.Request) {
	// Obtener los parámetros de la cadena de consulta usando request.URL.Query()
	vars := mux.Vars(request)

	// Establecer el encabezado "Content-Type" en la respuesta como "application/json"
	response.Header().Set("Content-Type", "application/json")

	// Agregar un encabezado personalizado llamado "Tamila" a la respuesta
	response.Header().Add("Tamila", "www.tamila.co")

	// Crear una estructura ResponseGenerico con el estado "ok" y un mensaje que incluye el ID de la solicitud GET
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Get Query String | id=" + vars["id"]})

	// Escribir la respuesta JSON en el cuerpo de la respuesta
	fmt.Fprintf(response, string(output))
}

func Ejemplo_get_con_parametros(response http.ResponseWriter, request *http.Request) {
	// Obtener los parámetros de la URL usando mux.Vars
	vars := mux.Vars(request)

	// Establecer el encabezado "Content-Type" en la respuesta como "application/json"
	response.Header().Set("Content-Type", "application/json")

	// Agregar un encabezado personalizado llamado "Tamila" a la respuesta
	response.Header().Add("Tamila", "www.tamila.co")

	// Crear una estructura ResponseGenerico con el estado "ok" y un mensaje que incluye el ID de la solicitud GET
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Get con Parametros | id=" + vars["id"]})

	// Escribir la respuesta JSON en el cuerpo de la respuesta
	fmt.Fprintf(response, string(output))
}

/*func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json") // Headers que se devuelven en
	response.Header().Add("Tamila", "www.tamila.co")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Post "})
	fmt.Fprintf(response, string(output))
}*/

/*// funcion que retorna un map en json y un status
func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json") // Headers que se devuelven en
	response.Header().Add("Tamila", "www.tamila.co")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Metodo POST2",
	}
	//response.WriteHeader(201)
	response.WriteHeader(http.StatusCreated) // sele da el status 201
	json.NewEncoder(response).Encode(respuesta)
}*/

// recibiendo datros desde postman
func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	// Establecer el encabezado "Content-Type" en la respuesta como "application/json"
	response.Header().Set("Content-Type", "application/json")

	// Crear una variable para almacenar la estructura de datos que se espera en el cuerpo JSON de la solicitud
	var categoria dto.CategoriaDTO
	// Decodificar el cuerpo JSON de la solicitud en la estructura de datos definida
	err := json.NewDecoder(request.Body).Decode(&categoria)

	// Verificar si hubo un error al decodificar el JSON
	if err != nil {
		log.Println("Error al decodificar JSON")
		http.Error(response, "Error al decodificar JSON:", http.StatusBadRequest)
		return
	}

	// Crear una respuesta JSON que incluye el estado "ok", un mensaje, y algunos datos de la estructura decodificada
	respuesta := map[string]string{
		"estado":        "ok",
		"mensaje":       "Metodo POST2",
		"nombre":        categoria.Nombre,
		"Authorizacion": request.Header.Get("Authorizacion"), // captura cabeceros
	}

	// Configurar la respuesta HTTP con el código 201 (Created)
	response.WriteHeader(http.StatusCreated) // sele da el status 201

	// Codificar la respuesta JSON y enviarla al cliente
	json.NewEncoder(response).Encode(respuesta)
}

// recibir datos desde la url
func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)                                 // captura el dato enviado en la url
	response.Header().Set("Content-Type", "application/json") // Headers que se devuelven en
	response.Header().Add("Tamila", "www.tamila.co")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Put | id=" + vars["id"] + "Nombre=" + vars["nombre"]})
	fmt.Fprintf(response, string(output))
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	// Obtener los parámetros de la URL usando mux.Vars
	vars := mux.Vars(request)

	// Establecer el encabezado "Content-Type" en la respuesta como "application/json"
	response.Header().Set("Content-Type", "application/json") // Headers que se devuelven en

	// Agregar un encabezado personalizado llamado "Tamila"
	response.Header().Add("Tamila", "www.tamila.co")

	// Crear una estructura ResponseGenerico con el estado "ok" y un mensaje que incluye el ID de la solicitud DELETE
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo Delete| id=" + vars["id"]})

	// Escribir la respuesta JSON en el cuerpo de la respuesta
	fmt.Fprintf(response, string(output))
}

// subida de archivos
func Ejemplo_upload(response http.ResponseWriter, request *http.Request) {
	// Obtener el archivo de la solicitud
	file, handler, _ := request.FormFile("foto")

	// Obtener la extensión del archivo
	var extension = strings.Split(handler.Filename, ".")[1]

	// Crear un nombre único para la imagen basado en la hora actual y la extensión del archivo
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension

	// Definir la ruta del archivo en el servidor
	var archivo string = "public/uploads/fotos/" + foto

	// Abrir o crear el archivo en el servidor
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		http.Error(response, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	// Copiar el contenido del archivo de la solicitud al archivo en el servidor
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(response, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// Crear una respuesta JSON indicando que la carga fue exitosa
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Metodo POST2",
		"foto":    foto,
	}

	// Configurar la respuesta HTTP con el código 201 (Created)
	response.WriteHeader(http.StatusCreated)

	// Codificar la respuesta JSON y enviarla al cliente
	json.NewEncoder(response).Encode(respuesta)
}

func EjemploVerFoto(response http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")
	if len(file) < 1 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	OpenFile, err := os.Open("public/uploads/" + request.URL.Query().Get("folder") + "/" + file)
	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	_, err = io.Copy(response, OpenFile)
	if err != nil {
		http.Error(response, "Error al copiar el archivo", http.StatusBadRequest)
	}
}
