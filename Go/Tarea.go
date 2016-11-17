package main

import (
    "github.com/go-martini/martini"
    //"github.com/martini-contrib/binding"
    "github.com/martini-contrib/render"
    "net/http"
    "io/ioutil"
    "./pkg/emails"
    "./pkg/upload"
    //"github.com/martini-contrib/cors"
    "fmt"
)



func main(){

    m := martini.Classic();

    m.Use(render.Renderer())

    m.Get("/",  func(r render.Render) {
        r.JSON(http.StatusOK, "Lenguajes de Programacion - Tareas - Nexer Rodriguez - 21411072")
    })

    m.Post("/orderEmails", func(r *http.Request, res render.Render){
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

    m.Post("/hideMessage", func(r *http.Request, res render.Render){
        err := r.ParseMultipartForm(100000)
        if err != nil {
            res.JSON(http.StatusInternalServerError, err.Error())
        }
        mensaje := r.FormValue("mensaje")
        imagenPath := "./uploads/img.bmp"
        upload.UploadFile(imagenPath, r)
        fmt.Println(mensaje)

        res.JSON(http.StatusOK, "Nothing yet")
    })

    m.Run()
}