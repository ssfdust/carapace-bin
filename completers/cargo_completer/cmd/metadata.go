package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Output the resolved dependencies of a package, the concrete used versions including overrides, in machine-readable format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metadataCmd).Standalone()

	metadataCmd.Flags().Bool("all-features", false, "Activate all available features")
	metadataCmd.Flags().StringSliceP("features", "F", []string{}, "Space or comma separated list of features to activate")
	metadataCmd.Flags().StringSlice("filter-platform", []string{}, "Only include resolve dependencies matching the given target-triple")
	metadataCmd.Flags().String("format-version", "", "Format version")
	metadataCmd.Flags().BoolP("help", "h", false, "Print help")
	metadataCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	metadataCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	metadataCmd.Flags().Bool("no-deps", false, "Output information only about the workspace members and don't fetch dependencies")
	metadataCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.AddCommand(metadataCmd)

	carapace.Gen(metadataCmd).FlagCompletion(carapace.ActionMap{
		"features":       action.ActionFeatures(metadataCmd).UniqueList(","),
		"format-version": carapace.ActionValues("1"),
		"manifest-path":  carapace.ActionFiles(),
	})
}
