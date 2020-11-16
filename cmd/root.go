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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/k1LoW/keyp/backend"
	"github.com/k1LoW/keyp/version"
	"github.com/spf13/cobra"
)

var (
	b        string
	users    []string
	groups   []string
	teams    []string
	keepKeys []string
	logTo    string
	create   bool
)

var rootCmd = &cobra.Command{
	Use:          "keyp",
	Short:        "keyp is a tool to keep public keys up to date",
	Long:         `keyp is a tool to keep public keys up to date.`,
	SilenceUsage: true,
	Version:      version.Version,
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
		if logTo == "" {
			debug, err := os.Create(fmt.Sprintf("%s.debug", version.Name))
			if err != nil {
				rootCmd.PrintErrln(err)
				os.Exit(1)
			}
			log.SetOutput(debug)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {}

func keys(ctx context.Context) ([]string, error) {
	client, err := backend.New(ctx, b)
	if err != nil {
		return nil, err
	}
	opts := []backend.Option{}
	if opt, err := backend.Users(users); err == nil {
		opts = append(opts, opt)
	} else {
		return nil, err
	}
	if b == "github" {
		if opt, err := backend.Teams(teams); err == nil {
			opts = append(opts, opt)
		} else {
			return nil, err
		}
	} else {
		if opt, err := backend.Groups(groups); err == nil {
			opts = append(opts, opt)
		} else {
			return nil, err
		}
	}

	keys, err := client.Keys(ctx, opts...)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, errors.New("key not found")
	}

	return keys, nil
}
