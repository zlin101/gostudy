package main

import (
	_ "gostudy/5-init/lib1"
	lib2 "gostudy/5-init/lib2"
)

func main() {
	// lib1.Lib1Init()
	lib2.Lib2Init()
	// lib1.Lib1Test()
	lib2.Lib2Test()
}
