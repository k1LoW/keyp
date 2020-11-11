/*
Copyright © 2020 Ken'ichiro Oyama <k1lowxb@gmail.com>

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
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/k1LoW/keyp/backend"
	"github.com/spf13/cobra"
)

// updateAuthorizedKeysCmd represents the updateAuthorizedKeys command
var updateAuthorizedKeysCmd = &cobra.Command{
	Use:   "update-authorized-keys [USER]",
	Short: "update [USER_HOME_DIR]/.ssh/authorized_keys",
	Long:  `update [USER_HOME_DIR]/.ssh/authorized_keys.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case logTo == "stdout":
			log.SetOutput(os.Stdout)
		case logTo == "stderr":
			log.SetOutput(os.Stderr)
		case logTo != "":
			l, err := os.OpenFile(filepath.Clean(logTo), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) // #nosec
			if err != nil {
				return err
			}
			log.SetOutput(l)
		}
		ctx := context.Background()
		keys, err := keys(ctx)
		if err != nil {
			return err
		}
		u, err := user.Lookup(args[0])
		if err != nil {
			return err
		}
		if u.HomeDir == "" {
			return fmt.Errorf("'%s' does not have home directory", u.Name)
		}
		aKeys := filepath.Join(u.HomeDir, ".ssh", "authorized_keys")
		if _, err := os.Stat(aKeys); err != nil {
			return err
		}
		return ioutil.WriteFile(aKeys, []byte(fmt.Sprintf("%s\n", strings.Join(keys, "\n"))), 0600)
	},
}

func init() {
	updateAuthorizedKeysCmd.Flags().StringVarP(&b, "backend", "b", "", fmt.Sprintf("backend key store %s (requied)", backend.Backends))
	updateAuthorizedKeysCmd.Flags().StringSliceVarP(&users, "user", "u", []string{}, "target user")
	updateAuthorizedKeysCmd.Flags().StringSliceVarP(&groups, "group", "g", []string{}, "target group")
	updateAuthorizedKeysCmd.Flags().StringSliceVarP(&teams, "team", "t", []string{}, "target org team")
	updateAuthorizedKeysCmd.Flags().StringVarP(&logTo, "log", "l", "", "log")
	if err := updateAuthorizedKeysCmd.MarkFlagRequired("backend"); err != nil {
		updateAuthorizedKeysCmd.PrintErrln(err)
		os.Exit(1)
	}
	rootCmd.AddCommand(updateAuthorizedKeysCmd)
}
