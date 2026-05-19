package cmd

import (
	"fmt"
	"os"

	"github.com/GuiCezaF/gocat/internal/reader"
	"github.com/spf13/cobra"
)

// countLinesCmd represents the countLines command
var countLinesCmd = &cobra.Command{
	Use:   "countLines [file_path]",
	Short: "Count lines in a file",
	Long:  `Analyze a file and return the total number of lines, useful for quick inspection, statistics, and file analysis directly from the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		file_path := args[0]
		file := reader.NewFile(file_path)
		n_lines, err := file.CountLines()
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(n_lines)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(countLinesCmd)
}
