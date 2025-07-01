package autobuild

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// BuildConfig defines the structure of gpmbuild.json
type BuildConfig struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Dependencies []string `json:"dependencies"`
	Build        []string `json:"build"`
	Cleanup      []string `json:"cleanup"`
}

// RunCommand executes shell commands
func RunCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: autobuild <repo-name | file-path>")
		return
	}

	// Extract argument
	arg := os.Args[1]

	var buildFile string
	if strings.HasPrefix(arg, "/") || strings.HasPrefix(arg, "./") {
		// If it's a direct file path
		buildFile = arg + "/gpmbuild.json"
	} else {
		// If it's a repo name
		repoName := strings.TrimSuffix(arg, ".git")
		buildFile = "./" + repoName + "/gpmbuild.json"
	}

	// Check if gpmbuild.json exists
	if _, err := os.Stat(buildFile); os.IsNotExist(err) {
		fmt.Println("No gpmbuild.json found at", buildFile+". You will need to manually compile.")
		return
	}

	// Read and parse gpmbuild.json
	data, err := ioutil.ReadFile(buildFile)
	if err != nil {
		fmt.Println("Error reading gpmbuild.json:", err)
		return
	}

	var config BuildConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing gpmbuild.json:", err)
		return
	}

	fmt.Printf("🚀 Building %s v%s...\n", config.Name, config.Version)

	// Install dependencies
	for _, dep := range config.Dependencies {
		fmt.Println("Installing dependency:", dep)
		RunCommand("apt-get", "install", "-y", dep)
	}

	// Execute build commands
	for _, cmd := range config.Build {
		fmt.Println("Executing:", cmd)
		RunCommand("bash", "-c", cmd)
	}

	// Cleanup steps
	for _, cmd := range config.Cleanup {
		fmt.Println("Cleaning up:", cmd)
		RunCommand("bash", "-c", cmd)
	}

	fmt.Println("✅ Build process complete!")
}
