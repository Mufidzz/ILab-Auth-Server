package main

import (
	"infotech.umm.ac.id/auth/config"
	"infotech.umm.ac.id/auth/data"
	//"github.com/Mufidzz/iflogs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.ClientDBInit()
	db := config.DBInit()
	InDB := data.InDB{DB: db}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "POST"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
	router.Use(ClientVerificator())
	router.POST("/", InDB.AuthorizeUser)
	router.GET("/clear", InDB.UnauthorizeUser)

	err := router.Run(":6666")
	//err := router.RunTLS(":6666", os.Getenv("TLS_CERT_FILE"), os.Getenv("TLS_KEY_FILE"))
	if err != nil {
		panic(err)
	}
}

func ClientVerificator() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.GetHeader("IFX-CLIENT")
		sec := c.GetHeader("IFX-SECRET")

		if cid == "" || sec == "" {
			c.JSON(http.StatusForbidden, "Unauthorized Client : 1")
			c.Abort()
			return
		}

		db := config.ClientDBInit()
		InDB := data.InDB{DB: db}

		cl, err := InDB.GetClientUsingBarrierIDWorker(cid)
		if err != nil {
			c.JSON(http.StatusForbidden, "Unauthorized Client : 2")
			c.Abort()
			return
		}

		if cl.BarrierSecret != sec {
			c.JSON(http.StatusForbidden, "Unauthorized Client : 3")
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
