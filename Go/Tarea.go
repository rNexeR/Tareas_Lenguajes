package main

import (
	"./pkg/emails"
	"./pkg/kruskal"
	"./pkg/steganography"
	"./pkg/upload"
	"encoding/json"
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
	//"github.com/martini-contrib/cors"
	"fmt"
)

func main() {

	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(http.StatusOK, "Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072")
	})

	m.Post("/orderEmails", func(r *http.Request, res render.Render) {
		err := r.ParseMultipartForm(100000)
		if err != nil {
			res.JSON(http.StatusInternalServerError, err.Error())
		}
		upload.UploadFile("./uploads/emails.txt", r)
		//call function to order emails
		filename := emails.OrderEmails("./uploads/emails.txt")
		retornar, _ := ioutil.ReadFile(filename)
		res.Data(http.StatusOK, retornar)

	})

	m.Post("/hideMessage", func(r *http.Request, res render.Render) {
		err := r.ParseMultipartForm(100000)
		if err != nil {
			res.JSON(http.StatusInternalServerError, err.Error())
		}
		mensaje := r.FormValue("mensaje")
		imagenPath := "./uploads/img.bmp"
		upload.UploadFile(imagenPath, r)
		fmt.Println(mensaje)

		steganography.WriteMessage(mensaje, imagenPath)

		retornar, _ := ioutil.ReadFile(imagenPath)
		res.Data(http.StatusOK, retornar)
	})

	m.Post("/discoverMessage", func(r *http.Request, res render.Render) {
		err := r.ParseMultipartForm(100000)
		if err != nil {
			res.JSON(http.StatusInternalServerError, err.Error())
		}
		imagenPath := "./uploads/img.bmp"
		upload.UploadFile(imagenPath, r)

		mensaje := steganography.ReadMessage(imagenPath)

		res.JSON(http.StatusOK, mensaje)
	})

	m.Post("/kruskal", func(r *http.Request, res render.Render) {
		r.ParseMultipartForm(100000)
		/*if err != nil {
			res.JSON(http.StatusInternalServerError, err.Error())
		}*/
		graphStr := r.FormValue("graph")
		var graph = kruskal.Graph{}
		json.Unmarshal([]byte(graphStr), &graph)
		kruskal.PrintGraph(graph)
		retornar := kruskal.Kruskal(graph)
		res.JSON(http.StatusOK, retornar)
	})

	m.Run()
}
