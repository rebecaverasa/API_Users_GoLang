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


func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil { 
		return 
	} 

	users = append(users, newUser)              
	c.IndentedJSON(http.StatusCreated, newUser)

}

func getUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getUserById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserById(id string) (*user, error) {
	for i, u := range users {
		if u.Id == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found")
}

func deleteUserById(c *gin.Context) {
	id := c.Param("id")

	for i, a := range users {
		if a.Id == id {
			aux := users[i+1:]
			users = append(users[:i], aux...)

			c.IndentedJSON(http.StatusOK, users)
			return
		}
	}
}


func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.PATCH("/users/:id", getUser)
	router.POST("/users", createUser)
	router.DELETE("/users/:id", deleteUserById)
	router.Run("localhost:9090")
}
