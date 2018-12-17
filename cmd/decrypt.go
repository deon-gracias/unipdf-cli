/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unidoc/unicli/pdf"
)

const decryptCmdDesc = ``

var decryptCmdExample = fmt.Sprintf("%s\n%s\n",
	fmt.Sprintf("%s decrypt -p pass input_file.pdf", appName),
	fmt.Sprintf("%s decrypt -p pass -o output_file.pdf input_file.pdf", appName),
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:                   "decrypt [FLAG]... INPUT_FILE",
	Short:                 "Decrypt PDF files",
	Long:                  decryptCmdDesc,
	Example:               decryptCmdExample,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		password, _ := cmd.Flags().GetString("password")

		// Parse output path.
		outputPath, _ := cmd.Flags().GetString("output-file")
		if outputPath == "" {
			outputPath = inputPath
		}

		// Decrypt input file.
		if err := pdf.Decrypt(inputPath, outputPath, password); err != nil {
			printErr("Could not decrypt input file: %s\n", err)
		}

		fmt.Printf("Successfully decrypted %s\n", inputPath)
		fmt.Printf("Output file saved to %s\n", outputPath)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Must provide the PDF file to decrypt\n")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringP("password", "p", "", "PDF file password")
	decryptCmd.Flags().StringP("output-file", "o", "", "Output file")
}
