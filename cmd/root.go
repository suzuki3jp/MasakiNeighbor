/*
Copyright © 2025 suzuki3jp
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/suzuki3jp/mn/internal/fs"
)

var rootCmd = &cobra.Command{
	Use:   "mn",
	Short: "最近隣空間的随伴尺度を求めるツール",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		points, err := fs.ReadPointsCsv(filename)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		cmd.Printf("読み込んだ点: %+v\n", points)
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
