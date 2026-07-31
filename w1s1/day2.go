package main

import "fmt"

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
	
	// 规则：
	// 1. 状态码必须是 100～599
	// 2. path 不能为空
	// 3. path 必须以 "/" 开头
	// 合法时返回 true, "ok"
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