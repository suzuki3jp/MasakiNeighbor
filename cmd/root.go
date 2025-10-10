/*
Copyright © 2025 suzuki3jp
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suzuki3jp/mn/internal/calc"
	"github.com/suzuki3jp/mn/internal/fs"
	"github.com/suzuki3jp/mn/internal/output"
)

var A float64

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

		result := calc.Mn(points, A)
		err = output.WriteJSONToFile(result, "./output.json")
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Float64VarP(&A, "a", "a", 1.0, "Parameter 'a' to pass to calc.Mn")
}
