/*
Copyright © 2025 @mdxabu

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "Sicko mode to show the visualization waves of the audio outputs.",
	Long: `Sicko mode is a command that visualizes the audio outputs in a wave format.`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)
}
