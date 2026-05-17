package cmd

import (
	"fmt"
  "os"

  "github.com/GuiCezaF/gocat/internal/reader"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read [file_path]",
	Short: "Read and display file contents",
	Long: `Read and print the contents of a file directly in the terminal with optional formatting features suchas has line numbering and syntax highlighting.`,
	Run: func(cmd *cobra.Command, args []string) {
    file_path := args[0]
    data, err := reader.ReadFile(file_path)
  
   if err != nil {
      os.Exit(1)
    }
    fmt.Println(string(data))
	},
  DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(readCmd)
}
