package main

import (
	"fmt"
	"os"

	"github.com/loveRyujin/ResuGo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "执行命令时出错: %v\n", err)
		os.Exit(1)
	}
}
