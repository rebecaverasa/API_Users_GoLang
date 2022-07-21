package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Userstatus int    `json:"userstatus"`
}

var users = []user{
	{Id: "1",
		Username:   "rebecaveras",
		Firstname:  "rebeca",
		Lastname:   "veras",
		Email:      "rebecaverasa@gmail.com",
		Password:   "senha",
		Phone:      "98959698",
		Userstatus: 1},

	{Id: "2",
		Username:   "joaosilva",
		Firstname:  "joao",
		Lastname:   "silva",
		Email:      "joaosilva@gmail.com",
		Password:   "senhaa",
		Phone:      "98959697",
		Userstatus: 2},

	{Id: "3",
		Username:   "mariosilva",
		Firstname:  "mario",
		Lastname:   "silva",
		Email:      "mariosilva@gmail.com",
		Password:   "teste",
		Phone:      "985748574",
		Userstatus: 3},
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var newUser user

	err := context.BindJSON(&newUser); 
	if err != nil {
		return
	}

	users = append(users, newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUser(context *gin.Context) {
	id := context.Param("id")      //pega o id que digitei na url
	user, err := getUserById((id)) //é o parametro id criado acima. E Acontece uma atribuição dupla: User e err pegam respectivamente 2 parametros da funçao getusersbyid: user e error

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user) //retorna o user do getuserbyid
}

func getUserById(id string) (*user, error) { //2 retornos: tipo user e tipo error.
	for i, u := range users { //range percorre todo o users. u é 1 elemento de users que modifica a cada interação do for.
		if u.Id == id { //u.Id pega só o Id do users e compara com o parametro de entrada. Ex:/2
			return &users[i], nil //o u vai procurar em looping o Id de entrada até encontrar.Quando achar, vai retnornar o usuario nesta posição e erro nil.
		}
	}

	return nil, errors.New("User not found") //quando percorrer tudo e nao encontrar o parametro de entrada, retornara um usuario nil e a mensagem de erro.
}

func deleteUserById(context *gin.Context) {
	id := context.Param("id") //parametro id da request delete.

	for i, a := range users {
		if a.Id == id {
			aux := users[i+1:]              //aqui pega os usuarios depois do usuario que digitei
			users = append(users[:i], aux...) //o users deixa de ser o elemento completo e passa a ser o users sem o numero que quero deletar

			context.IndentedJSON(http.StatusOK, users)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.DELETE("/users/:id", deleteUserById)
	router.Run("localhost:9090")
}
