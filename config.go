package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Config struct {
	Author     string `json:"author"`
	Since      string `json:"since"`
	ShowMerges bool   `json:"show-merges"`
	ShowDate   bool   `json:"show-date"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".config", "whatdidido")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

func loadConfig() Config {
	defaultAuthor := getGitUserName()
	defaultSince := "1 week ago"
	defaultShowMerges := false
	defaultShowDate := true

	configPath, err := getConfigPath()
	if err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince, ShowMerges: defaultShowMerges, ShowDate: defaultShowDate}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince, ShowMerges: defaultShowMerges, ShowDate: defaultShowDate}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince, ShowMerges: defaultShowMerges, ShowDate: defaultShowDate}
	}

	// Ensure defaults if fields are empty
	if config.Author == "" {
		config.Author = defaultAuthor
	}
	if config.Since == "" {
		config.Since = defaultSince
	}
	// ShowMerges defaults to false, no need to set
	// ShowDate defaults to true for new configs, but respect saved value

	return config
}

func saveConfig(config Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func handleConfigCommand(cmd *cobra.Command) {
	authorFlag, err := cmd.Flags().GetString("author")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting author flag: %v\n", err)
		os.Exit(1)
	}
	sinceFlag, err := cmd.Flags().GetString("since")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting since flag: %v\n", err)
		os.Exit(1)
	}
	showMergesFlagStr, err := cmd.Flags().GetString("show-merges")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting show merges flag: %v\n", err)
		os.Exit(1)
	}

	// check if flag is present
	if !cmd.Flags().Changed("author") && !cmd.Flags().Changed("since") && !cmd.Flags().Changed("show-merges") && !cmd.Flags().Changed("show-date") {
		// Show current config
		config := loadConfig()
		fmt.Printf("Author: %s\n", config.Author)
		fmt.Printf("Since:  %s\n", config.Since)
		fmt.Printf("Show Merges: %t\n", config.ShowMerges)
		fmt.Printf("Show Date: %t\n", config.ShowDate)
		return
	}

	config := loadConfig()
	if authorFlag != "" {
		config.Author = authorFlag
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Author set to: %s\n", authorFlag)
	}
	if sinceFlag != "" {
		config.Since = sinceFlag
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Since set to: %s\n", sinceFlag)
	}
	if cmd.Flags().Changed("show-merges") {
		showMergesFlag := showMergesFlagStr == "true"
		config.ShowMerges = showMergesFlag
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Show Merges set to: %t\n", showMergesFlag)
	}

	showDateFlagStr, _ := cmd.Flags().GetString("show-date")
	if cmd.Flags().Changed("show-date") {
		showDateFlag := showDateFlagStr == "true"
		config.ShowDate = showDateFlag
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Show Date set to: %t\n", showDateFlag)
	}
}
