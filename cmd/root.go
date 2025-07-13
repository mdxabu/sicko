/*
Copyright © 2025 @mdxabu

*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"fmt"
)



var rootCmd = &cobra.Command{
	Use:   "sicko",
	Short: "Sicko was the CLI Audio Visualizer made for fun :)",
	Long: `Sicko is a CLI Audio Visualizer that was created for fun and experimentation.`,

	Run: func(cmd *cobra.Command, args []string) {
		asciiBanner := `
  ██████  ██▓ ▄████▄   ██ ▄█▀ ▒█████  
▒██    ▒ ▓██▒▒██▀ ▀█   ██▄█▒ ▒██▒  ██▒
░ ▓██▄   ▒██▒▒▓█    ▄ ▓███▄░ ▒██░  ██▒
  ▒   ██▒░██░▒▓▓▄ ▄██▒▓██ █▄ ▒██   ██░
▒██████▒▒░██░▒ ▓███▀ ░▒██▒ █▄░ ████▓▒░
▒ ▒▓▒ ▒ ░░▓  ░ ░▒ ▒  ░▒ ▒▒ ▓▒░ ▒░▒░▒░ 
░ ░▒  ░ ░ ▒ ░  ░  ▒   ░ ░▒ ▒░  ░ ▒ ▒░ 
░  ░  ░   ▒ ░░        ░ ░░ ░ ░ ░ ░ ▒  
      ░   ░  ░ ░      ░  ░       ░ ░  
             ░                        		
		`
		fmt.Println(asciiBanner)
	},

}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


