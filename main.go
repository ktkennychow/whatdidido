package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whatdidido",
	Short: "A simple CLI tool to show you easy to read git commits history.",
	Long:  `A simple CLI tool to show you easy to read git commits history. Or you are just too lazy to memorize git flags.`,
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check git commits",
	Run: func(cmd *cobra.Command, args []string) {
		handleCheckCommand(cmd)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Run: func(cmd *cobra.Command, args []string) {
		handleConfigCommand(cmd)
	},
}

func init() {
	checkCmd.Flags().StringP("author", "a", "", "Git author name (e.g., \"johndoe\")")
	checkCmd.Flags().StringP("since", "s", "", "Time specification (e.g., \"1 day ago\", \"midnight\")")
	checkCmd.Flags().StringP("show-merges", "m", "", "Show merge commits")
	checkCmd.Flags().StringP("show-date", "d", "", "Show date before commit message (default: true)")

	configCmd.Flags().StringP("author", "a", "", "Git author name (e.g., \"johndoe\")")
	configCmd.Flags().StringP("since", "s", "", "Time specification (e.g., \"1 day ago\", \"midnight\")")
	configCmd.Flags().StringP("show-merges", "m", "", "Show merge commits")
	configCmd.Flags().StringP("show-date", "d", "", "Show date before commit message (default: true)")

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(configCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		// Cobra handles errors
	}
}
