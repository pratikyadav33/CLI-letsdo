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

type todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done int    `json:"done"`
	Date string `json:"date"`
}

func check() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("Select * From todo;")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ID \t \t Task \t \t \t \t Status? \t \t Time Created  ")
	fmt.Println("")

	for results.Next() {
		var td todo
		err = results.Scan(&td.ID, &td.Task, &td.Done, &td.Date)
		if err != nil {
			panic(err.Error())
		}
		fmt.Print(td.ID)
		fmt.Print("\t \t" + td.Task + "\t \t")
		if td.Done == 0 {
			fmt.Print("Not Completed")
		} else {
			fmt.Print("Completed")
		}

		fmt.Print("\t \t ")
		fmt.Print(td.Date)
		fmt.Println("")

	}

}
func checkCompleted() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("Select * From todo where done = 1;")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ID \t \t Task \t \t \t \t Status? \t \t Time Created  ")
	fmt.Println("")

	for results.Next() {
		var td todo
		err = results.Scan(&td.ID, &td.Task, &td.Done, &td.Date)
		if err != nil {
			panic(err.Error())
		}
		fmt.Print(td.ID)
		fmt.Print("\t \t" + td.Task + "\t \t")
		if td.Done == 0 {
			fmt.Print("Not Completed")
		} else {
			fmt.Print("Completed")
		}

		fmt.Print("\t \t ")
		fmt.Print(td.Date)
		fmt.Println("")

	}
}

func checkNotCompleted() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("Select * From todo where done = 0;")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ID \t \t Task \t \t \t \t Status? \t \t Time Created  ")
	fmt.Println("")
	for results.Next() {
		var td todo
		err = results.Scan(&td.ID, &td.Task, &td.Done, &td.Date)
		if err != nil {
			panic(err.Error())
		}
		fmt.Print(td.ID)
		fmt.Print("\t \t" + td.Task + "\t \t")
		if td.Done == 0 {
			fmt.Print("Not Completed")
		} else {
			fmt.Print("Completed")
		}

		fmt.Print("\t \t ")
		fmt.Print(td.Date)
		fmt.Println("")

	}
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for task ",
	Long: `letsgo check can be used to keep track of completed and not completed Task Example: 
	letsdo check <flag -c(for completed task )> -u(for not completed task)`,
	Run: func(cmd *cobra.Command, args []string) {
		flagStatusC, _ := cmd.Flags().GetBool("completed")
		flagStatusN, _ := cmd.Flags().GetBool("not_completed")

		if flagStatusC {
			checkCompleted()
		} else if flagStatusN {
			checkNotCompleted()
		} else {
			check()
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolP("completed", "c", false, "Shows Completed Task Example: letsgo check -c")
	checkCmd.Flags().BoolP("not_completed", "u", false, "Shows Not Completed Task Example: letsgo check -u")

}
