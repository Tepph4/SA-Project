package main

import (
	"github.com/Tepph4/SA-Project/controller"
	"github.com/Tepph4/SA-Project/entity"
	"github.com/Tepph4/SA-Project/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// User Routes
			router.GET("/users", controller.ListUsers)
			router.GET("/user/:id", controller.GetUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			// Role Routes 
			r.GET("/roles", controller.ListRoles)
			r.GET("/role/:id", controller.GetRole)

			// Building Routes
			router.GET("/buildings", controller.ListBuildings)
			router.GET("/building/:id", controller.GetBuilding)
			router.POST("/buildings", controller.CreateBuilding)
			router.PATCH("/buildings", controller.UpdateBuilding)
			router.DELETE("/buildings/:id", controller.DeleteBuilding)

			// Room Routes
			router.GET("/rooms", controller.ListRooms)
			router.GET("/room/:id", controller.GetRoom)
			router.GET("/room/building/:id", controller.GetRoomByBuilding)
			router.POST("/rooms", controller.CreateRoom)
			router.PATCH("/rooms", controller.UpdateRoom)
			router.DELETE("/rooms/:id", controller.DeleteRoom)

			// Device Routes
			router.GET("/devices", controller.ListDevices)
			router.GET("/device/:id", controller.GetDevice)
			router.POST("/devices", controller.CreateDevice)
			router.PATCH("/devices", controller.UpdateDevice)
			router.DELETE("/devices/:id", controller.DeleteDevice)

			// Room has Device Routes
			router.GET("/room_has_devices", controller.ListRoomHasDevice)
			router.GET("/roomhasdevice/:id", controller.GetRoomHasDevice)
			router.POST("/room_has_devices", controller.CreateRoomHasDevice)
			router.PATCH("/room_has_devices", controller.UpdateRoomHasDevice)
			router.DELETE("/roomhasdevices/:id", controller.DeleteRoomHasDevice)

		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateUser)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
