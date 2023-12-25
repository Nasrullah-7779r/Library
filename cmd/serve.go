package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"library/pkg/student"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `Applcation will be served on host and port defined in config.yml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	router := gin.Default()

	student.SetupStudentRoutes(router)

	router.Run()
}
