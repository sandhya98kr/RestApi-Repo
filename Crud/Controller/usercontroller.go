package controller

import (
	"Crud/database"
	"fmt"

	"Crud/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	//automigrate() will create tables,fix missing indexes.
	db.AutoMigrate(&entity.User{})
	return &UserRepo{Db: db}
}

//create user
//gin context contains http.request and http/response
func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	err := entity.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println("user created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

//get users
func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []entity.User
	err := entity.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"message": "listed all the users in table", "list": user})
}

//get user by id
func (repository *UserRepo) GetUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user entity.User
	err := entity.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

// update user
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user entity.User
	id, _ := c.Params.Get("id")
	err := entity.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println("user updated successfully")
	c.BindJSON(&user)

	err = entity.UpdateUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// delete user
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user entity.User
	id, _ := c.Params.Get("id")
	err := entity.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
