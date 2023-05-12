package routes

import (
	controller "github.com/davidrogger/first-golang-crud/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:id", controller.GetUserByID)
	r.GET("/getUserByEmail/:email", controller.GetUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)
}
