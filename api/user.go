package api

import (
	"net/http"

	database "garrison/internal/database/postgres"
	"garrison/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []*models.User

	database.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}
