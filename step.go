package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunEnvmanAdd(key, value string) error {
	args := []string{"add", "-k", key, "-v", value}
	return RunCommand("envman", args...)
}
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// unix timestamp
	timestamp := time.Now().Unix()
	timestampString := fmt.Sprintf("%d", timestamp)
	if err := RunEnvmanAdd("UNIX_TIMESTAMP", timestampString); err != nil {
		fmt.Println("Failed to store UNIX_TIMESTAMP:", err)
		os.Exit(1)
	}

	// iso8601 time format (timezone: RFC3339Nano)
	timeString := fmt.Sprintf("%v", time.Now().Format(time.RFC3339Nano))
	if err := RunEnvmanAdd("ISO_DATETIME", timeString); err != nil {
		fmt.Println("Failed to store ISO_DATETIME:", err)
		os.Exit(1)
	}
}
