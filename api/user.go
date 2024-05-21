package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:id`
	Name  string `json:name`
	Email string `json:email`
}

var users = []User{
	{ID: "1", Name: "Felipe Heredia de Moura", Email: "hello@felipeheredia.com"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
