package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Author string `json:"author"`
	Since  string `json:"since"`
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
	defaultSince := "midnight"

	configPath, err := getConfigPath()
	if err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{Author: defaultAuthor, Since: defaultSince}
	}

	// Ensure defaults if fields are empty
	if config.Author == "" {
		config.Author = defaultAuthor
	}
	if config.Since == "" {
		config.Since = defaultSince
	}

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

func handleConfigCommand() {
	var authorFlag string
	var sinceFlag string
	flag.StringVar(&authorFlag, "author", "", "Git author name (e.g., \"johndoe\")")
	flag.StringVar(&authorFlag, "a", "", "Git author name (shorthand for --author)")
	flag.StringVar(&sinceFlag, "since", "", "Time specification (e.g., \"1 day ago\", \"midnight\")")
	flag.StringVar(&sinceFlag, "s", "", "Time specification (shorthand for --since)")
	flag.Parse()

	if authorFlag == "" && sinceFlag == "" {
		// Show current config
		config := loadConfig()
		fmt.Printf("Author: %s\n", config.Author)
		fmt.Printf("Since:  %s\n", config.Since)
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
}
