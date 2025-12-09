package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	Author string `json:"author"`
	Since  string `json:"since"`
}

func getGitUserName() string {
	cmd := exec.Command("git", "config", "user.name")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
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
	if len(os.Args) == 2 {
		// Show current config
		config := loadConfig()
		fmt.Printf("Author: %s\n", config.Author)
		fmt.Printf("Since:  %s\n", config.Since)
		return
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage:")
		fmt.Println("  whatdidido config                    - Show current config")
		fmt.Println("  whatdidido config author <name>      - Set author")
		fmt.Println("  whatdidido config since <time>        - Set since")
		os.Exit(1)
	}

	config := loadConfig()
	key := os.Args[2]
	value := os.Args[3]

	switch key {
	case "author":
		config.Author = value
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Author set to: %s\n", value)
	case "since":
		config.Since = value
		if err := saveConfig(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Since set to: %s\n", value)
	default:
		fmt.Fprintf(os.Stderr, "Unknown config key: %s\n", key)
		fmt.Println("Valid keys: author, since")
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("whatdidido - A simple CLI tool to show your git commits since midnight.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  whatdidido [author] [since]")
	fmt.Println("  whatdidido --author=<name> | -a <name>")
	fmt.Println("  whatdidido --since=<time> | -s <time>")
	fmt.Println("  whatdidido config [key] [value]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("  author  Git author name (default: git config user.name, or empty)")
	fmt.Println("  since   Time specification (default: midnight)")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --author, -a  Git author name (e.g., \"johndoe\")")
	fmt.Println("  --since, -s   Time specification (e.g., \"1 day ago\", \"midnight\")")
	fmt.Println()
	fmt.Println("Config:")
	fmt.Println("  whatdidido config                    - Show current config")
	fmt.Println("  whatdidido config author <name>      - Set author")
	fmt.Println("  whatdidido config since <time>       - Set since")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  whatdidido")
	fmt.Println("  whatdidido johndoe")
	fmt.Println("  whatdidido johndoe \"1 day ago\"")
	fmt.Println("  whatdidido --author johndoe")
	fmt.Println("  whatdidido -a johndoe --since \"1 day ago\"")
	fmt.Println("  whatdidido --since \"1 day ago\"")
	fmt.Println("  whatdidido -s \"1 day ago\"")
	fmt.Println("  whatdidido config author johndoe")
	fmt.Println("  whatdidido config since \"1 day ago\"")
	fmt.Println()
	fmt.Printf("This runs: git log --since=<since> --author=<author> --no-merges --pretty=format:\"%%s\"\n")
}

func main() {
	// Check for help flags
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "-h" || arg == "--help" || arg == "help" {
			showHelp()
			return
		}
		if arg == "config" {
			handleConfigCommand()
			return
		}
	}

	// Parse flags
	var authorFlag string
	var sinceFlag string
	flag.StringVar(&authorFlag, "author", "", "Git author name (e.g., \"johndoe\")")
	flag.StringVar(&authorFlag, "a", "", "Git author name (shorthand for --author)")
	flag.StringVar(&sinceFlag, "since", "", "Time specification (e.g., \"1 day ago\", \"midnight\")")
	flag.StringVar(&sinceFlag, "s", "", "Time specification (shorthand for --since)")
	flag.Parse()

	// Load config as defaults
	config := loadConfig()
	author := config.Author
	since := config.Since

	// Override with flags if provided
	if authorFlag != "" {
		author = authorFlag
	}
	if sinceFlag != "" {
		since = sinceFlag
	}

	// Handle positional arguments (author and optionally since if not set via flags)
	posArgs := flag.Args()
	if len(posArgs) > 0 && authorFlag == "" {
		author = posArgs[0]
	}
	if len(posArgs) > 1 && sinceFlag == "" {
		since = posArgs[1]
	}

	args := []string{"log",
		fmt.Sprintf("--since=%s", since),
		"--no-merges",
		`--pretty=format:%s`,
	}

	// Only add --author flag if author is not empty
	if author != "" {
		args = append(args, fmt.Sprintf("--author=%s", author))
	}

	cmd := exec.Command("git", args...)

	output, err := cmd.Output()
	if err != nil {
		// git log returns exit code 1 when no commits found, which is normal
		if len(output) == 0 {
			fmt.Println("no commit found - someone is slacking!")
			return
		}
		fmt.Fprintf(os.Stderr, "Error running git log: %v\n", err)
		os.Exit(1)
	}

	if len(output) == 0 {
		fmt.Println("no commit found - someone is slacking!")
		return
	}

	fmt.Print(string(output))
}
