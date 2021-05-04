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

func update() {
	fmt.Println("Check Usage")
}
func updateDelete() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP Table todo;")
	if err != nil {
		panic(err.Error())
	}

}

func updateTask(args []string) {
	var tempTask string
	var tempid string
	for _, val := range args {
		temp := val
		tempTask = tempTask + temp + " "
	}

	tempid = tempTask[:1]
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Query("UPDATE `todo` SET `task` = '" + tempTask[1:] + "' WHERE `todo`.`id` = " + tempid + ";")
	if err != nil {
		panic(err.Error())
	}
	check()
}
func updateTaskDone(args []string) {
	var tempid string
	for _, val := range args {
		temp := val
		tempid = tempid + temp
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var ch string
	ch = tempid[1:]
	ch = ch[:1]
	var st string

	if ch == "1" {
		st = "1"
	} else {
		st = "0"
	}
	//UPDATE `todo` SET `done` = '1' WHERE `todo`.`id` = " + temp_id + ";
	_, err = db.Query("UPDATE `todo` SET `done` = " + st + " WHERE `todo`.`id` = " + tempid[:1] + ";")
	if err != nil {
		panic(err.Error())
	}
	check()

}

func updateStask(args []string) {
	var tempTask string
	for _, val := range args {
		temp := val
		tempTask = tempTask + temp
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("Delete from todo where id = " + tempTask + ";")
	if err != nil {
		panic(err.Error())
	}
	check()
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Task",
	Long:  `letsdo update can be used to keep mordify database and task  Task Example: letsgo update --help`,
	Run: func(cmd *cobra.Command, args []string) {
		flagStatusdb, _ := cmd.Flags().GetBool("deletedatabase")
		flagStatust, _ := cmd.Flags().GetBool("task")
		flagStatusd, _ := cmd.Flags().GetBool("taskdone")
		flagStatuss, _ := cmd.Flags().GetBool("deletetask")

		if flagStatusd {
			updateTaskDone(args)

		} else if flagStatust {
			updateTask(args)

		} else if flagStatusdb {
			updateDelete()
		} else if flagStatuss {
			updateStask(args)
		} else {
			update()
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolP("deletedatabase", "a", false, "Delete all Task")
	updateCmd.Flags().BoolP("deletetask", "d", false, "Delete Task by ID Example letsdo update <Taskid> -d")
	updateCmd.Flags().BoolP("task", "t", false, "Change the Name of Task Example: letsdo update <Task Id> <New Name> -t")
	updateCmd.Flags().BoolP("taskdone", "s", false, "Change if task is done or not Example: letsdo update <Task Id> <Value 1(Completed) or 0(Not )>  -s")
}
