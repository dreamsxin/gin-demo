/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gin-demo/middleware"
	"gin-demo/router"

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

var debug bool

func init() {
	rootCmd.AddCommand(serveCmd)
	viper.SetDefault("addr", ":8080")
}

func run() {
	r := gin.Default()
	middleware.InitMiddleware(r, viper.GetViper())
	router.InitRouter(r, viper.GetViper())
	r.Run(viper.GetString("addr")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
