package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:id", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)
}
