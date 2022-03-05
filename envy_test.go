package envy

import (
	"testing"
)

func TestLoad(t *testing.T) {
	port := Load("PORT")
	if port == "" {
		t.Errorf("Error in TestingLoad\n")
	}
}

func TestEnvironments(t *testing.T) {
	// Define environments
	newEnvs := map[string]string{
		"production": ".env",
		"staging":    "file.json",
	}

	// Setting Environments
	SetEnvironments(newEnvs)

	// Set Current Environment
	SetActiveEnv("production")

	// Load PORT Key
	portVal := Load("PORT")
	if portVal != "5000" {
		t.Error("Load 1 - Not compatible value")
	}

	// Change Environment
	SetActiveEnv("staging")

	// Load keyTwo Key
	keyTwo := Load("keyTwo")
	if keyTwo != "valueTwo" {
		t.Error("Load 2 - Not compatible value")
	}

	// Reset environments
	Reset()

	// Load PORT from default environment file
	portVal = Load("PORT")
	if portVal != "5000" {
		t.Error("Load 3 - Not compatible value")
	}
}
