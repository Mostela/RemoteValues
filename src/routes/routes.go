package routes

import (
	"RemoteValues/src/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func HandlerRequest() {
	r := gin.Default()
	r.GET("/healthcheck", controllers.HealthCheckControl)
	r.GET("/keyrotation", controllers.ReturnKeysRotation)
	r.POST("/keyrotation", controllers.SetNewKeyRotation)
	r.NoRoute(controllers.Return404)
	log.Fatal(r.Run())
}
