package v1

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jegacs/golang-bootcamp-2020/handlers/v1/countries"

	cors "github.com/itsjamie/gin-cors"
)

var (
	Router *gin.Engine
	v1     *gin.RouterGroup
)

func init() {
	Router = gin.New()
	Router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/"))
	Router.Use(gin.Recovery())

	Router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,DELETE,PUT",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 = Router.Group("v1")

	countries.Countries(v1)
}

func Run(addr ...string) error {
	return Router.Run(addr...)
}
