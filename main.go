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
package main

import (
	"database/sql"

	"letsdo/cmd"

	_ "github.com/go-sql-driver/mysql"
)

func createAndOpen(name string) *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	return db
}
func createTable(table string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/letsdo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `letsdo`.`todo` ( `id` INT NOT NULL AUTO_INCREMENT , `task` VARCHAR(1000) NOT NULL , `done` BOOLEAN NOT NULL , `date` DATETIME NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}
	db.Close()

}
func main() {

	db := createAndOpen("letsdo")
	createTable("todo")
	defer db.Close()
	cmd.Execute()
}
