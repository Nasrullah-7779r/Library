package cmd

import (
	"github.com/spf13/cobra"
	"library/database"
)

func init() {
	rootCmd.AddCommand(migrateCMD)
}

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Tables to the DB",
	Long:  `Migrate Tables to the DB`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {

	database.Set()
	database.Connect()
	database.Migration()
}
