package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func parseGapiFile(filePath string, visited map[string]bool) error {
    if visited[filePath] {
        return nil // prevent recursive includes
    }
    visited[filePath] = true

    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open %s: %v", filePath, err)
    }
    defer file.Close()

    var section string
    var repos []string
    var autobuildDir, outputDir string
    logFile := false

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
            section = line
            continue
        }

        switch section {
        case "[REPO_LINKS]":
            repos = append(repos, line)
        case "[AUTOBUILD_DIR]":
            autobuildDir = line
        case "[OUTPUT_DIR]":
            outputDir = line
        case "[LOG_FILE]":
            logFile = strings.ToLower(line) == "true"
        case "[INCLUDE]":
            if err := parseGapiFile(line, visited); err != nil {
                return err
            }
        }
    }

    for _, repo := range repos {
        fmt.Printf("Cloning %s...\n", repo)
        cmd := exec.Command("git", "clone", repo, outputDir)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
            fmt.Printf("Error cloning %s: %v\n", repo, err)
        }
    }

    if autobuildDir != "" {
        fmt.Println("Running autobuild...")
        cmd := exec.Command(autobuildDir, outputDir)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
            fmt.Printf("Error running autobuild: %v\n", err)
        }
    }

    if logFile {
        fmt.Println("Logging enabled (mock behavior here).")
    }

    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: gpm-api <file.gapi>")
        return
    }

    err := parseGapiFile(os.Args[1], map[string]bool{})
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
