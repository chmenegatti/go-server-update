package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-server-updater/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:id", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser/", userController.CreateUser)
	r.PUT("/updateUser/:id", userController.UpdateUser)
	r.DELETE("/deleteUser/:id", userController.DeleteUser)
}
