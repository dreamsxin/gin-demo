/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gin-demo/config"
	"gin-demo/debug"
	"gin-demo/middleware"
	"gin-demo/router"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Application",
	Long:  "Starts the application and listens for incoming requests",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	viper.SetDefault("addr", ":8080")
}

func run() {
	r := gin.Default()
	middleware.InitMiddleware(r, viper.GetViper())
	router.InitRouter(r, viper.GetViper())
	log.Println("debug", isDebug, viper.GetBool("debug"))
	if isDebug {
		debug.InitDebug()
	}

	docroot := config.Server().WebServer.Root
	if docroot == "" {
		docroot = "./"
	}
	fs := static.LocalFile(docroot, false)
	fileserver := http.FileServer(fs)

	r.NoRoute(func(c *gin.Context) {
		prefix := "/"
		if fs.Exists(prefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
	r.Run(viper.GetString("addr")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
