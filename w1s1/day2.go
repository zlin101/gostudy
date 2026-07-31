package main

import (
	"fmt"
	"strings"
	"strconv"
)

func classifyStatus(status int) (string, bool) {
	switch {
	case status >= 200 && status < 300:
		return "success", true
	case status >= 300 && status < 400:
		return "redirect", true
	case status >= 400 && status < 600:
		return "error", true
	default:
		return "unknown", false
	}
}

func validateLogLine(line string) (bool, string) {
	fields := strings.Fields(line)

	if len(fields) != 2 {
		return false, "invalid log line"
	}

	status, err := strconv.Atoi(fields[0])
	if err != nil {
		return false, "invalid status"
	}

	return validateLog(status, fields[1])
}

func validateLog(status int, path string) (bool, string) {
	if status < 100 || status >= 600 {
		return false, "invalid status"
	}

	if len(path) == 0 {
		return false, "empty path"
	}

	if !strings.HasPrefix(path, "/") {
		return false, "invalid path"
	}

	return true, "ok"
}

/*
200 /users true ok
700 /admin false invalid status
404  false empty path
201 orders false invalid path
*/

func main() {

	lines := []string{
		"200 /users",
		"abc /users",
		"200",
		"",
		"404 /health",
		"700 /admin",
		"201 orders",
		"200 /users extra",
	}
	var validCount int

	for _, line := range lines {
		valid, reason := validateLogLine(line)
		fmt.Printf("%q -> %t, %s\n", line, valid, reason)

		if valid {
			validCount++
		}
	}
		fmt.Printf("valid count: %d\n", validCount)
}