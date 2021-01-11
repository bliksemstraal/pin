/*
Copyright © 2021 bliksemstraal (bliksemstraal12@outlook.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/bliksemstraal/pin/password"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pin",
	Short: "Generate secure passwords from a key and passphrase",
	Long: `
                _
   *      _ __ (_)_ __
7  8  9  | '_ \| | '_ \
4  5  6  | |_) | | | | |
1  2  3  | .__/|_|_| |_|
   0     |_|
	
Pin generates strong, unique passwords for web accounts without storing them on your
device.
   `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("\nPlease provide a key word to know what this password is for")
			return cmd.Help()
		}

		secret, err := typePassphrase()
		if err != nil {
			return err
		}

		input := args[0] + "??!!??" + secret

		gen := password.New(24, input)
		fmt.Println(gen.Encrypt())
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetUsageFunc(func(f *cobra.Command) error {
		fmt.Println("Usage:")
		fmt.Println("  pin KEYWORD")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  To (re)generate a password for you Google account:*")
		fmt.Println("  pin google.com")
		fmt.Println()
		fmt.Println("* NOTE: you will be prompted for a secret twice (use the same secret for all passwords)")
		fmt.Println()
		return nil
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func typePassphrase() (string, error) {
	var (
		p1, p2 []byte
		err    error
		count  int
	)
	for count < 3 {
		fmt.Println("Enter secret: ")
		p1, err = terminal.ReadPassword(0)
		if err != nil {
			return "", err
		}
		fmt.Println("Confirm secret: ")
		p2, err = terminal.ReadPassword(0)
		if err != nil {
			return "", err
		}
		if string(p1) == string(p2) {
			return string(p1), nil
		}
		fmt.Println("\nTry again")
	}
	return "", errors.New("passwords did not match")
}
