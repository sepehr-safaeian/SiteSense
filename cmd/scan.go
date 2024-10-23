package cmd

import (
	"SiteSense/internal"

	"github.com/spf13/cobra"
)

func ScanPage(cmd *cobra.Command, args []string) {
	URL := args[0]

	internal.Collector(URL)
}
