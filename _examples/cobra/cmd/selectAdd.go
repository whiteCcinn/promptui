/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

/*
➜  cobra git:(master) ✗ go run main.go selectAdd
Other: add
✔ add
You choose add
selectAdd called
 */

// selectAddCmd represents the selectAdd command
var selectAddCmd = &cobra.Command{
	Use:   "selectAdd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		items := []string{"Vim", "Emacs", "Sublime", "VSCode", "Atom"}
		index := -1
		var result string
		var err error

		for index < 0 {
			prompt := promptui.SelectWithAdd{
				Label:    "What's your text editor",
				Items:    items,
				AddLabel: "Other",
			}

			index, result, err = prompt.Run()

			if index == -1 {
				items = append(items, result)
			}
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %s\n", result)

		fmt.Println("selectAdd called")
	},
}

func init() {
	rootCmd.AddCommand(selectAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}