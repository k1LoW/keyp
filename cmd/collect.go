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
	"os"

	"github.com/k1LoW/keyp/backend"
	"github.com/spf13/cobra"
)

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect keys from backend key store",
	Long:  `collect keys from backend key store.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		keys, err := keys(ctx)
		if err != nil {
			return err
		}
		for _, k := range keys {
			cmd.Println(k)
		}
		return nil
	},
}

func init() {
	collectCmd.Flags().StringVarP(&b, "backend", "b", "", fmt.Sprintf("backend key store %s (requied)", backend.Backends))
	collectCmd.Flags().StringSliceVarP(&users, "user", "u", []string{}, "target user")
	collectCmd.Flags().StringSliceVarP(&groups, "group", "g", []string{}, "target group")
	collectCmd.Flags().StringSliceVarP(&teams, "team", "t", []string{}, "target org team")
	if err := collectCmd.MarkFlagRequired("backend"); err != nil {
		collectCmd.PrintErrln(err)
		os.Exit(1)
	}
	rootCmd.AddCommand(collectCmd)
}
