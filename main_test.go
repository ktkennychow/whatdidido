package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGetConfigPath tests the config file path generation
func TestGetConfigPath(t *testing.T) {
	path, err := getConfigPath()
	if err != nil {
		t.Fatalf("getConfigPath failed: %v", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get home dir: %v", err)
	}

	expected := filepath.Join(homeDir, ".config", "whatdidido", "config.json")
	if path != expected {
		t.Errorf("Expected path %s, got %s", expected, path)
	}
}

// TestLoadConfig_NoConfigFile tests loading config when no config file exists
func TestLoadConfig_NoConfigFile(t *testing.T) {
	// Remove any existing config file first
	configPath, _ := getConfigPath()
	os.Remove(configPath)

	config := loadConfig()

	// Should have defaults
	if config.Author == "" {
		t.Error("Expected non-empty default author")
	}
	if config.Since == "" {
		t.Error("Expected non-empty default since")
	}
	if config.ShowMerges != false {
		t.Error("Expected ShowMerges to default to false")
	}
	if config.ShowDate != false {
		t.Error("Expected ShowDate to default to false")
	}
}

// TestSaveLoadConfig tests saving and loading config
func TestSaveLoadConfig(t *testing.T) {
	testConfig := Config{
		Author:     "Test Author",
		Since:      "3 days ago",
		ShowMerges: true,
		ShowDate:   true,
	}

	err := saveConfig(testConfig)
	if err != nil {
		t.Fatalf("saveConfig failed: %v", err)
	}

	loadedConfig := loadConfig()
	if loadedConfig.Author != testConfig.Author {
		t.Errorf("Author mismatch: expected %s, got %s", testConfig.Author, loadedConfig.Author)
	}
	if loadedConfig.Since != testConfig.Since {
		t.Errorf("Since mismatch: expected %s, got %s", testConfig.Since, loadedConfig.Since)
	}
	if loadedConfig.ShowMerges != testConfig.ShowMerges {
		t.Errorf("ShowMerges mismatch: expected %t, got %t", testConfig.ShowMerges, loadedConfig.ShowMerges)
	}
	if loadedConfig.ShowDate != testConfig.ShowDate {
		t.Errorf("ShowDate mismatch: expected %t, got %t", testConfig.ShowDate, loadedConfig.ShowDate)
	}
}

// TestLoadConfig_EmptyFields tests loading config with empty fields (should use defaults)
func TestLoadConfig_EmptyFields(t *testing.T) {
	// Save config with empty fields
	emptyConfig := Config{
		Author:     "",
		Since:      "",
		ShowMerges: false,
		ShowDate:   true,
	}

	err := saveConfig(emptyConfig)
	if err != nil {
		t.Fatalf("saveConfig failed: %v", err)
	}

	loadedConfig := loadConfig()

	// Should fill in defaults for empty fields
	if loadedConfig.Author == "" {
		t.Error("Expected default author when field is empty")
	}
	if loadedConfig.Since == "" {
		t.Error("Expected default since when field is empty")
	}
}

// TestGetGitUserName tests getting git user name
func TestGetGitUserName(t *testing.T) {
	name := getGitUserName()

	// Should return a string (might be empty if no git config)
	if name != "" {
		// If not empty, should look like a name (contains space or is reasonable length)
		if !strings.Contains(name, " ") && len(name) < 2 {
			t.Errorf("Git user name looks invalid: %s", name)
		}
	}
	t.Logf("Git user name: %s", name)
}

// TestConfigStruct tests Config struct initialization
func TestConfigStruct(t *testing.T) {
	config := Config{
		Author:     "test@example.com",
		Since:      "1 week ago",
		ShowMerges: true,
		ShowDate:   true,
	}

	if config.Author != "test@example.com" {
		t.Errorf("Author not set correctly")
	}
	if config.Since != "1 week ago" {
		t.Errorf("Since not set correctly")
	}
	if config.ShowMerges != true {
		t.Errorf("ShowMerges not set correctly")
	}
	if config.ShowDate != true {
		t.Errorf("ShowDate not set correctly")
	}
}

// TestSaveConfig_ValidData tests that saveConfig works with valid data
func TestSaveConfig_ValidData(t *testing.T) {
	validConfig := Config{
		Author:     "valid@example.com",
		Since:      "1 day ago",
		ShowMerges: false,
		ShowDate:   true,
	}

	err := saveConfig(validConfig)
	if err != nil {
		t.Errorf("saveConfig should not fail with valid config: %v", err)
	}
}

// TestConfigWithEmptyAuthor tests config handling with empty author flag
func TestConfigWithEmptyAuthor(t *testing.T) {
	// Set a known config first
	originalConfig := Config{
		Author:     "Original Author",
		Since:      "1 week ago",
		ShowMerges: false,
		ShowDate:   true,
	}
	err := saveConfig(originalConfig)
	if err != nil {
		t.Fatalf("Failed to save original config: %v", err)
	}

	// Load config - should preserve original author
	loadedConfig := loadConfig()
	if loadedConfig.Author != "Original Author" {
		t.Errorf("Expected original author to be preserved, got: %s", loadedConfig.Author)
	}
}

// TestConfigWithEmptySince tests config handling with empty since flag
func TestConfigWithEmptySince(t *testing.T) {
	// Set a known config first
	originalConfig := Config{
		Author:     "Test Author",
		Since:      "2 weeks ago",
		ShowMerges: false,
		ShowDate:   true,
	}
	err := saveConfig(originalConfig)
	if err != nil {
		t.Fatalf("Failed to save original config: %v", err)
	}

	// Load config - should preserve original since
	loadedConfig := loadConfig()
	if loadedConfig.Since != "2 weeks ago" {
		t.Errorf("Expected original since to be preserved, got: %s", loadedConfig.Since)
	}
}

// TestConfigWithShowMerges tests config handling with show-merges flag
func TestConfigWithShowMerges(t *testing.T) {
	testConfig := Config{
		Author:     "Test Author",
		Since:      "1 week ago",
		ShowMerges: true,
		ShowDate:   true,
	}

	err := saveConfig(testConfig)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	loadedConfig := loadConfig()
	if loadedConfig.ShowMerges != true {
		t.Errorf("Expected ShowMerges to be true, got: %t", loadedConfig.ShowMerges)
	}
}

// TestConfigWithShowDate tests config handling with show-date flag
func TestConfigWithShowDate(t *testing.T) {
	// Test with ShowDate = true
	testConfig := Config{
		Author:     "Test Author",
		Since:      "1 week ago",
		ShowMerges: false,
		ShowDate:   true,
	}

	err := saveConfig(testConfig)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	loadedConfig := loadConfig()
	if loadedConfig.ShowDate != true {
		t.Errorf("Expected ShowDate to be true, got: %t", loadedConfig.ShowDate)
	}

	// Test with ShowDate = false
	testConfig.ShowDate = false
	err = saveConfig(testConfig)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	loadedConfig = loadConfig()
	if loadedConfig.ShowDate != false {
		t.Errorf("Expected ShowDate to be false, got: %t", loadedConfig.ShowDate)
	}
}

// TestLoadConfig_DefaultShowDate tests that ShowDate defaults to false for new configs
func TestLoadConfig_DefaultShowDate(t *testing.T) {
	// Remove any existing config file first
	configPath, _ := getConfigPath()
	os.Remove(configPath)

	config := loadConfig()

	// ShowDate should default to false
	if config.ShowDate != false {
		t.Errorf("Expected ShowDate to default to false, got: %t", config.ShowDate)
	}
}

// TestSaveLoadConfig_AllFields tests saving and loading config with all fields
func TestSaveLoadConfig_AllFields(t *testing.T) {
	testConfig := Config{
		Author:     "Full Test Author",
		Since:      "5 days ago",
		ShowMerges: true,
		ShowDate:   false,
	}

	err := saveConfig(testConfig)
	if err != nil {
		t.Fatalf("saveConfig failed: %v", err)
	}

	loadedConfig := loadConfig()
	if loadedConfig.Author != testConfig.Author {
		t.Errorf("Author mismatch: expected %s, got %s", testConfig.Author, loadedConfig.Author)
	}
	if loadedConfig.Since != testConfig.Since {
		t.Errorf("Since mismatch: expected %s, got %s", testConfig.Since, loadedConfig.Since)
	}
	if loadedConfig.ShowMerges != testConfig.ShowMerges {
		t.Errorf("ShowMerges mismatch: expected %t, got %t", testConfig.ShowMerges, loadedConfig.ShowMerges)
	}
	if loadedConfig.ShowDate != testConfig.ShowDate {
		t.Errorf("ShowDate mismatch: expected %t, got %t", testConfig.ShowDate, loadedConfig.ShowDate)
	}
}

// TestConfigStruct_AllFields tests Config struct initialization with all fields
func TestConfigStruct_AllFields(t *testing.T) {
	config := Config{
		Author:     "test@example.com",
		Since:      "1 week ago",
		ShowMerges: true,
		ShowDate:   false,
	}

	if config.Author != "test@example.com" {
		t.Errorf("Author not set correctly")
	}
	if config.Since != "1 week ago" {
		t.Errorf("Since not set correctly")
	}
	if config.ShowMerges != true {
		t.Errorf("ShowMerges not set correctly")
	}
	if config.ShowDate != false {
		t.Errorf("ShowDate not set correctly")
	}
}

// TestConfigDefaultValues tests that default values are correct
func TestConfigDefaultValues(t *testing.T) {
	// Remove any existing config file first
	configPath, _ := getConfigPath()
	os.Remove(configPath)

	config := loadConfig()

	// Check all defaults
	if config.Author == "" {
		t.Error("Expected non-empty default author")
	}
	if config.Since != "1 week ago" {
		t.Errorf("Expected default since '1 week ago', got: %s", config.Since)
	}
	if config.ShowMerges != false {
		t.Error("Expected ShowMerges to default to false")
	}
	if config.ShowDate != false {
		t.Error("Expected ShowDate to default to false")
	}
}
