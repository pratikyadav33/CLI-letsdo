/*
Copyright © 2021 Pratik <PRATIKYADAV33@GMAIL.com>

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
	"fmt"

	"github.com/spf13/cobra"
)

func addTask(args []string) {
	var tempTask string
	for _, val := range args {
		temp := val
		tempTask = tempTask + temp + " "
	}
	if len(tempTask) >= 1 {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query("INSERT INTO todo(task,done) VALUES('" + tempTask + "', 0);")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Task Added : " + tempTask)
		defer insert.Close()

		check()

	} else {
		fmt.Println("Cant Leave the Task Empty ")
	}

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Task",
	Long: `letsdo add  -->  lets you add task to complete. For Example:
	letsdo add <Task-Name>`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
