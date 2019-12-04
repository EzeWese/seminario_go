package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", GetHello)

	if err := router.Run(); err != nil {
		fmt.Println(err)
	}

	router.GET("/cars/:plate", GetName)

}

func GetHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func GetName(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}

/*type Auto struct {
	marca   string
	modelo  string
	patente string
}

func NewAuto(patente string, marca string, modelo string) Auto {
	return Auto{
		patente: patente,
		marca:   marca,
		modelo:  modelo,
	}
}*/

/*type Agencia struct {
	modelo string
	autos  []Auto
}

func NewAgencia(nombreModelo string) Agencia {
	return Agencia{
		modelo: nombreModelo,
	}
}

func (a *Agencia) AgregarAuto(autoNuevo Auto) {
	a.autos = append(a.autos, autoNuevo)
}*/

/*func (a *Agencia) BorrarAuto(auto Auto) {
	return append
}*/

/*func main() {
	a := NewAgencia("Ford")
	fmt.Println("agencia creada")
	a.AgregarAuto(NewAuto("1234", "Ford", "4000"))
	fmt.Println(a.autos)

}*/

/*func main() {
	fmt.Printf("hello, world\n")
	c := otraCarpeta.SumarValores(5, 6)
	fmt.Println(c)
}*/
