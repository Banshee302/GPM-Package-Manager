package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// RunCommand executes shell commands
func RunCommand(cmd string, args ...string) error {
    command := exec.Command(cmd, args...)
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr
    return command.Run()
}

func convertToSSH(url string) string {
    // e.g. https://github.com/user/repo.git → git@github.com:user/repo.git
    url = strings.TrimPrefix(url, "https://")
    url = strings.TrimPrefix(url, "http://")
    parts := strings.SplitN(url, "/", 2)
    if len(parts) == 2 {
        return fmt.Sprintf("git@%s:%s", parts[0], parts[1])
    }
    return url // fallback
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage:\n  gpm -i <repo-url> [--ssh]\n  gpm -f <repo-url>\n  gpm -r <package-dir>")
        return
    }

    command := os.Args[1]
    argument := os.Args[2]
    useSSH := len(os.Args) > 3 && os.Args[3] == "--ssh"

    switch command {
    case "-i": // Install package
        cloneURL := argument
        if useSSH {
            cloneURL = convertToSSH(argument)
            fmt.Println("Using SSH:", cloneURL)
        } else {
            fmt.Println("Using HTTPS:", cloneURL)
        }

        fmt.Println("Cloning:", cloneURL)
        if err := RunCommand("git", "clone", cloneURL); err != nil {
            fmt.Println("Error cloning repository:", err)
            return
        }

        // Extract repo name
        repoName := strings.TrimSuffix(argument[strings.LastIndex(argument, "/")+1:], ".git")
        repoPath := "./" + repoName

        // Check for gpmbuild.json
        buildFile := repoPath + "/gpmbuild.json"
        if _, err := os.Stat(buildFile); os.IsNotExist(err) {
            fmt.Println("No gpmbuild.json found in", repoName+". You will need to manually compile.")
        } else {
            fmt.Println("Running autobuild for", repoName)
            if err := RunCommand("./autobuild/autobuild", repoPath); err != nil {
                fmt.Println("Error running autobuild:", err)
            }
        }

    case "-f": // Find repo
        fmt.Println("Searching:", argument)
        RunCommand("git", "ls-remote", argument)

    case "-r": // Remove package
        fmt.Println("Uninstalling:", argument)
        if err := os.RemoveAll(argument); err != nil {
            fmt.Println("Error removing package:", err)
        }

    default:
        fmt.Println("Unknown command:", command)
    }
}
