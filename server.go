package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/FxIvan/sisalcloud-contenedor-go/blob/main/excercise/twoParameterUrl"
)

//Definicion de structura de JSON
type album struct{
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64 		`json:"price"`
}

var albums = []album{
	{ID:"1", Title:"Without Me", Artist:"Eminem", Price: 10},
	{ID:"2", Title:"Normal", Artist:"Feid", Price: 9.99},
	{ID:"3", Title:"Fragil", Artist:"Yahritza y Su Esenncia", Price: 15},
	{ID:"4", Title:"Por ti", Artist:"Luana", Price: 8.99},
	{ID:"5", Title:"Gata Caliente", Artist:"LOLO OG y Omar varela", Price: 6.85},
	{ID:"6", Title:"TQG", Artist:"KAROL G Y Shakira", Price: 14.99},
}

func getAlbums(c *gin.Context){
	//Envia la respuesta en json
	c.IndentedJSON(http.StatusOK, albums)
}

func getIdAlbums(c *gin.Context){
	id := c.Param("id")

	for _,a := range albums{
		 if a.ID == id {
			c.IndentedJSON(http.StatusOK , a)
			return
		 }
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album not found"})
}

func postAlbums(c *gin.Context){

	var newAlbum album
	//console.log pero en go
	//obtenemos la direccion de la memoria de &newAlbum
	if err := c.BindJSON(&newAlbum);
	err != nil{
		return
	}
	albums = append(albums,newAlbum)
	fmt.Println(albums)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main(){
	//En este caso ponemos gin.Default() para que nos permite manejar las peticiones entrantes y decir que funcion tiene que hacer ante un peticion
	//Puede manejar GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id",getIdAlbums)
	router.GET("/twoparameter/:title/:page",twoparameter.twoParameter())

	//Corre el servidor en el puerto 8080
	router.Run("127.0.0.1:8080")
}