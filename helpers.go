package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getGitUserName() string {
	cmd := exec.Command("git", "config", "user.name")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
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
