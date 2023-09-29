package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilleLange/phoenix/pkg/phoenix"
	"github.com/spf13/cobra"
)

var (
	newLocation string
	projectPath string
	rootCmd     = &cobra.Command{
		Use:   "phoenix",
		Short: "Phoenix is an easy-to-use CLI to reborn Go projects into an other git repository.",
		Long:  "Phoenix is an easy-to-use CLI to reborn Go projects into an other git repository. It replace the given projet old location by the new one in all files.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := phoenix.Migrate(projectPath, newLocation); err != nil {
				log.Fatalln("can't migrate the project :", err.Error())
			}
			log.Println("The project is reborn!")
		},
	}
)

func Execute() {
	rootCmd.Flags().StringVarP(&newLocation, "location", "l", "github.com", "The new project location URL.")
	rootCmd.Flags().StringVarP(&projectPath, "path", "p", ".", "The project directory path.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
