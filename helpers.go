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
	fmt.Println("whatdidido - A simple CLI tool to show you easy to read git commits history.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  whatdidido check [--author=<name>] [--since=<time>]")
	fmt.Println("  whatdidido config [--author=<name>] [--since=<time>]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --author, -a  Git author name (e.g., \"johndoe\")")
	fmt.Println("  --since, -s   Time specification (e.g., \"1 day ago\", \"midnight\")")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  whatdidido check")
	fmt.Println("  whatdidido check --author johndoe")
	fmt.Println("  whatdidido check -a johndoe --since \"1 day ago\"")
	fmt.Println("  whatdidido check --since \"1 day ago\"")
	fmt.Println("  whatdidido check -s \"1 day ago\"")
	fmt.Println("  whatdidido config")
	fmt.Println("  whatdidido config --author johndoe")
	fmt.Println("  whatdidido config --since \"1 day ago\"")
	fmt.Println()
	fmt.Printf("This runs: git log --since=<since> --author=<author> --no-merges --pretty=format:\"%%s\"\n")
}
