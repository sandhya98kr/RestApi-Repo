package main

import (
	"Crud/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

//Engine is framework's instance, it contails middleware
func setupRouter() *gin.Engine {
	//default returns engine instance with logger and recovery middleware
	r := gin.Default()
	userRepo := controller.New()
	r.POST("/user", userRepo.CreateUser)
	r.GET("/user", userRepo.GetUsers)
	r.GET("/user/:id", userRepo.GetUser)
	r.PUT("/user/:id", userRepo.UpdateUser)
	r.DELETE("/user/:id", userRepo.DeleteUser)

	return r

}

//200 statusok
//404 page not found
//500 internal server error
