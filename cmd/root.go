package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Start",
	Short: "start using SiteSense",
	Run:   startApp,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startApp(cmd *cobra.Command, args []string) {
	fmt.Println("salam")
}
