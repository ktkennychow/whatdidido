package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

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
