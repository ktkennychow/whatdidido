package main

import (
	"os/exec"
	"strings"
	"testing"
)

// TestCheckCommand_NoFlags tests the check command without any flags
func TestCheckCommand_NoFlags(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test the check command
	cmd := exec.Command("./whatdidido-test", "check")
	output, err := cmd.Output()

	// The command might fail if no git repo, but we want to test it runs
	t.Logf("Command output: %s", string(output))
	if err != nil {
		t.Logf("Command failed as expected (no git repo): %v", err)
	}
}

// TestConfigCommand_ShowConfig tests the config command to show configuration
func TestConfigCommand_ShowConfig(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test the config command
	cmd := exec.Command("./whatdidido-test", "config")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config command failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Author:") {
		t.Errorf("Expected 'Author:' in output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Since:") {
		t.Errorf("Expected 'Since:' in output, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Show Merges:") {
		t.Errorf("Expected 'Show Merges:' in output, got: %s", outputStr)
	}
}

// TestConfigCommand_SetAuthor tests setting author via config command
func TestConfigCommand_SetAuthor(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test setting author
	cmd := exec.Command("./whatdidido-test", "config", "--author", "Test Author")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config set author failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Author set to: Test Author") {
		t.Errorf("Expected author set message, got: %s", outputStr)
	}
}

// TestConfigCommand_EmptyAuthorFlag tests setting empty author (should not change)
func TestConfigCommand_EmptyAuthorFlag(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// First set a known author
	setCmd := exec.Command("./whatdidido-test", "config", "--author", "Original Author")
	err = setCmd.Run()
	if err != nil {
		t.Fatalf("Failed to set original author: %v", err)
	}

	// Now try to set empty author - should not change
	emptyCmd := exec.Command("./whatdidido-test", "config", "--author", "")
	err = emptyCmd.Run()
	if err != nil {
		t.Fatalf("Empty author flag failed: %v", err)
	}

	// Check that config still has the original author
	showCmd := exec.Command("./whatdidido-test", "config")
	output, err := showCmd.Output()
	if err != nil {
		t.Fatalf("Show config failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Author: Original Author") {
		t.Errorf("Expected original author to remain, got: %s", outputStr)
	}
}

// TestInvalidCommand tests behavior with invalid command
func TestInvalidCommand(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test invalid command
	cmd := exec.Command("./whatdidido-test", "invalid-command")
	output, err := cmd.CombinedOutput()

	// Cobra handles unknown commands gracefully, so check for error message
	outputStr := string(output)
	if !strings.Contains(outputStr, "unknown command") {
		t.Errorf("Expected 'unknown command' in output, got: %s", outputStr)
	}
}

// TestCheckCommand_WithAuthorFlag tests check command with author flag
func TestCheckCommand_WithAuthorFlag(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test check command with author flag
	cmd := exec.Command("./whatdidido-test", "check", "--author", "Test User")
	_, err = cmd.Output()

	// Command might fail due to no git repo, but should not crash
	if err != nil {
		t.Logf("Command failed as expected (no git repo): %v", err)
	}
}

// TestCheckCommand_WithSinceFlag tests check command with since flag
func TestCheckCommand_WithSinceFlag(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test check command with since flag
	cmd := exec.Command("./whatdidido-test", "check", "--since", "1 day ago")
	_, err = cmd.Output()

	// Command might fail due to no git repo, but should not crash
	if err != nil {
		t.Logf("Command failed as expected (no git repo): %v", err)
	}
}

// TestCheckCommand_WithEmptySinceFlag tests check command with empty since flag
func TestCheckCommand_WithEmptySinceFlag(t *testing.T) {
	// Build the binary first
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test check command with empty since flag
	cmd := exec.Command("./whatdidido-test", "check", "--since", "")
	_, err = cmd.Output()

	// Command might fail due to no git repo, but should not crash
	if err != nil {
		t.Logf("Command failed as expected (no git repo): %v", err)
	}
}

// TestCheckCommand_WithShowMergesFlag tests check command with show-merges flag
func TestCheckCommand_WithShowMergesFlag(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test with show-merges=true
	cmd := exec.Command("./whatdidido-test", "check", "--show-merges", "true")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with show-merges=true: %v", err)
	}

	// Test with show-merges=false
	cmd = exec.Command("./whatdidido-test", "check", "--show-merges", "false")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with show-merges=false: %v", err)
	}
}

// TestCheckCommand_WithShowDateFlag tests check command with show-date flag
func TestCheckCommand_WithShowDateFlag(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test with show-date=true
	cmd := exec.Command("./whatdidido-test", "check", "--show-date", "true")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with show-date=true: %v", err)
	}

	// Test with show-date=false
	cmd = exec.Command("./whatdidido-test", "check", "--show-date", "false")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with show-date=false: %v", err)
	}
}

// TestCheckCommand_WithAllFlags tests check command with all flags combined
func TestCheckCommand_WithAllFlags(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	cmd := exec.Command("./whatdidido-test", "check",
		"--author", "Test User",
		"--since", "1 day ago",
		"--show-merges", "true",
		"--show-date", "false")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with all flags: %v", err)
	}
}

// TestCheckCommand_WithShortFlags tests check command with short flag aliases
func TestCheckCommand_WithShortFlags(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test short flags: -a (author), -s (since), -m (show-merges), -d (show-date)
	cmd := exec.Command("./whatdidido-test", "check", "-a", "Test", "-s", "1 day ago", "-m", "false", "-d", "true")
	_, err = cmd.Output()
	if err != nil {
		t.Logf("Command with short flags: %v", err)
	}
}

// TestConfigCommand_SetShowMerges tests setting show-merges via config command
func TestConfigCommand_SetShowMerges(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test setting show-merges to true
	cmd := exec.Command("./whatdidido-test", "config", "--show-merges", "true")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config set show-merges failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Show Merges set to: true") {
		t.Errorf("Expected show-merges set message, got: %s", outputStr)
	}

	// Verify it was saved
	showCmd := exec.Command("./whatdidido-test", "config")
	output, _ = showCmd.Output()
	if !strings.Contains(string(output), "Show Merges: true") {
		t.Errorf("Expected Show Merges: true in config, got: %s", string(output))
	}

	// Test setting show-merges to false
	cmd = exec.Command("./whatdidido-test", "config", "--show-merges", "false")
	output, err = cmd.Output()
	if err != nil {
		t.Fatalf("Config set show-merges to false failed: %v", err)
	}

	outputStr = string(output)
	if !strings.Contains(outputStr, "Show Merges set to: false") {
		t.Errorf("Expected show-merges set message, got: %s", outputStr)
	}
}

// TestConfigCommand_SetShowDate tests setting show-date via config command
func TestConfigCommand_SetShowDate(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	// Test setting show-date to true
	cmd := exec.Command("./whatdidido-test", "config", "--show-date", "true")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config set show-date failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Show Date set to: true") {
		t.Errorf("Expected show-date set message, got: %s", outputStr)
	}

	// Verify it was saved
	showCmd := exec.Command("./whatdidido-test", "config")
	output, _ = showCmd.Output()
	if !strings.Contains(string(output), "Show Date: true") {
		t.Errorf("Expected Show Date: true in config, got: %s", string(output))
	}

	// Test setting show-date to false
	cmd = exec.Command("./whatdidido-test", "config", "--show-date", "false")
	output, err = cmd.Output()
	if err != nil {
		t.Fatalf("Config set show-date to false failed: %v", err)
	}

	outputStr = string(output)
	if !strings.Contains(outputStr, "Show Date set to: false") {
		t.Errorf("Expected show-date set message, got: %s", outputStr)
	}
}

// TestConfigCommand_SetSince tests setting since via config command
func TestConfigCommand_SetSince(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	cmd := exec.Command("./whatdidido-test", "config", "--since", "3 days ago")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config set since failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Since set to: 3 days ago") {
		t.Errorf("Expected since set message, got: %s", outputStr)
	}

	// Verify it was saved
	showCmd := exec.Command("./whatdidido-test", "config")
	output, _ = showCmd.Output()
	if !strings.Contains(string(output), "Since:  3 days ago") {
		t.Errorf("Expected Since: 3 days ago in config, got: %s", string(output))
	}
}

// TestConfigCommand_SetMultipleFlags tests setting multiple flags at once
func TestConfigCommand_SetMultipleFlags(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	cmd := exec.Command("./whatdidido-test", "config",
		"--author", "Multi Test",
		"--since", "2 days ago",
		"--show-merges", "true",
		"--show-date", "false")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config set multiple flags failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Author set to: Multi Test") {
		t.Errorf("Expected author set message, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Since set to: 2 days ago") {
		t.Errorf("Expected since set message, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Show Merges set to: true") {
		t.Errorf("Expected show-merges set message, got: %s", outputStr)
	}
	if !strings.Contains(outputStr, "Show Date set to: false") {
		t.Errorf("Expected show-date set message, got: %s", outputStr)
	}
}

// TestConfigCommand_ShowsShowDate tests that config command shows Show Date field
func TestConfigCommand_ShowsShowDate(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "whatdidido-test")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer exec.Command("rm", "whatdidido-test").Run()

	cmd := exec.Command("./whatdidido-test", "config")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Config command failed: %v", err)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "Show Date:") {
		t.Errorf("Expected 'Show Date:' in output, got: %s", outputStr)
	}
}
