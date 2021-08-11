// @title Swagger Example api
//@version 2.0
// @description This is a stand alone server
//@contact.email sandhya.kr@quest-global.com
// @host localhost:8080
// @BasePath /user
// @Success 200 {array} entity.User <-- This is a user defined struct.
// @Success 404 {array} entity.User
// @Success 500 {array} entity.User

package main

import (
	"Crud/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger/example/basic/docs"
	//"github.com/alecthomas/template"

	"github.com/gin-gonic/gin"
)

func main() {

	r := setupRouter()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	_ = r.Run(":8080")
}

//Engine is framework's instance, it contails middleware
func setupRouter() *gin.Engine {
	//default returns engine instance with logger and recovery middleware
	r := gin.Default()
	userRepo := controller.New()
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.POST("/user", userRepo.CreateUser)
	r.GET("/user", userRepo.GetUsers)
	r.GET("/user/:id", userRepo.GetUser)
	r.PUT("/user/:id", userRepo.UpdateUser)
	r.DELETE("/user/:id", userRepo.DeleteUser)

	return r

}

//404 page not found
//500 internal server error
