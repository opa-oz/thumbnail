package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/thumbnail/pkg/image"
	"github.com/spf13/cobra"
)

func newFname(oldFname string) string {
	filename := filepath.Base(oldFname)
	filenameParts := strings.Split(filename, ".")

	return strings.Replace(oldFname, filename, fmt.Sprintf("%s_thumbnail.%s", filenameParts[0], filenameParts[1]), 1)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thumbnail",
	Short: "Generate thumbnail from image",
	Long:  `Generate thumbnail from input image`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Please, specify input image")
			os.Exit(1)
		}

		length := todo.Int("Do thumbnail for multiple input files", 1)
		width := todo.Int("Width as argument", 160)
		height := todo.Int("Height as argument", 224)

		for i := 0; i < length; i++ {
			filename := args[0]

			err := image.ProcessImage(filename, uint(width), uint(height), newFname)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
