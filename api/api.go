package api

import (
	"github.com/MarcosOliveir4/api-lol/api/routes"
	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
