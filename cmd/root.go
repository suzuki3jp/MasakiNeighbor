/*
Copyright © 2025 suzuki3jp
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mn",
	Short: "最近隣空間的随伴尺度を求めるツール",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, World!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
