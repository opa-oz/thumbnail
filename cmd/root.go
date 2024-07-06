package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/thumbnail/pkg/image"
	"github.com/opa-oz/thumbnail/pkg/validators"
	"github.com/spf13/cobra"
)

func newFname(oldFname string) string {
	filename := filepath.Base(oldFname)
	filenameParts := strings.Split(filename, ".")

	return strings.Replace(oldFname, filename, fmt.Sprintf("%s_thumbnail.%s", filenameParts[0], filenameParts[1]), 1)
}

var size string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thumbnail",
	Short: "Generate thumbnail from image",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		for _, arg := range args {
			if err := validators.ExistsOrError(arg); err != nil {
				return err
			}

			filename := filepath.Base(arg)
			filenameParts := strings.Split(filename, ".")

			if err := validators.ExtensionOrError(&filenameParts, filename); err != nil {
				return err
			}

			extension := filenameParts[1]

			if err := validators.SupportedOrError(extension); err != nil {
				return err
			}
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validators.ValidateSize(size)
		if err != nil {
			return err
		}

		parts := strings.Split(size, "x")
		width, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}

		height, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		length := todo.Int("Do thumbnail for multiple input files", 1)

		for i := 0; i < length; i++ {
			filename := args[0]

			err := image.ProcessImage(filename, uint(width), uint(height), newFname)

			if err != nil {
				return err
			}
		}

		return nil
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
	rootCmd.Flags().StringVarP(&size, "size", "s", "160x224", "Size of produced thumbnail - <width>x<height>")
	rootCmd.MarkFlagRequired("size")
}
