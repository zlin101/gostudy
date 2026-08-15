//go:build ignore

package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Name string `json:"name"`
	Year int    `json:"year"`
	Actor []string `json:"actor"`
}

func main() {
	var m Movie
	m.Name = "重庆森林"
	m.Year = 1994
	m.Actor = []string{"王菲", "梁朝伟", "林青霞"}

	// 将结构体转换为json字符串
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println("jsonStr=", string(jsonStr))

	// 将json字符串转换为结构体
	var m2 Movie
	err = json.Unmarshal(jsonStr, &m2)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	fmt.Println("m2=", m2)
}
