package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Work with the system cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(cacheCmd)
}