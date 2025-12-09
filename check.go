package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func handleCheckCommand(cmd *cobra.Command) {
	authorFlag, _ := cmd.Flags().GetString("author")
	sinceFlag, _ := cmd.Flags().GetString("since")

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

	args := []string{"log",
		fmt.Sprintf("--since=%s", since),
		"--no-merges",
		`--pretty=format:%s`,
	}

	// Only add --author flag if author is not empty
	if author != "" {
		args = append(args, fmt.Sprintf("--author=%s", author))
	}

	gitCmd := exec.Command("git", args...)

	output, err := gitCmd.Output()
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
