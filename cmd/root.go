package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	DeepScan bool
	URL      string
)

var rootCmd = &cobra.Command{
	Use:   "sitesense",
	Short: "Start using SiteSense",
	Run:   startApp,
}

var scanPageCmd = &cobra.Command{
	Use:   "scan [URL]",
	Short: "Start scanning a page with SiteSense",
	Long:  `This command scans the provided web page URL and returns performance, SEO, and other metrics.`,
	Args:  cobra.ExactArgs(1),
	Run:   ScanPage,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startApp(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to SiteSense CLI")
}

func init() {
	rootCmd.AddCommand(scanPageCmd)

	scanPageCmd.Flags().BoolVarP(&DeepScan, "deep", "d", false, "Perform a deep scan")
}
