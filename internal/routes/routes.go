package routes

import (
	"github.com/SssHhhAaaDddOoWww/Goiler/internal/routes/health"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/health", health.Health)
}
