package main

import {
	"fmt"
	"string"
}

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

func main() {
	statuses := []int{200, 302, 707, 404, 700}
	validCount := 0

	for i, status := range statuses {
		kind, ok := classifyStatus(status)

		if !ok {
			fmt.Println(i, status, "invalid")
			continue
		}

		if kind == "success" {
			validCount++
		}

		fmt.Println(i, status, kind)
	}

	fmt.Println("validCount:", validCount)
}