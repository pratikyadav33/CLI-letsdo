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
	"database/sql"

	"github.com/spf13/cobra"
)

func done(args []string) {
	var temp_id string

	for _, ival := range args {
		itemp := ival
		temp_id = temp_id + itemp
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Query("UPDATE `todo` SET `done` = '1' WHERE `todo`.`id` = " + temp_id + ";")
	if err != nil {
		panic(err.Error())
	}
	check()

}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark Task Done ",
	Long: `For example: 
	letsdo done <ID>`,
	Run: func(cmd *cobra.Command, args []string) {
		done(args)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
