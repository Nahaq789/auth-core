package presentation

import (
	"github.com/gin-gonic/gin"
)

func CreateRoot(r gin.IRouter) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup")
	}
}
