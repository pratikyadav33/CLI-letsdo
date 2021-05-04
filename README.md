# CLI-letsdo

# letsdo 
## cli todo manager you need

letsdo is a cli, platform-independent todo manager

## Features

- Add task directly from cli interface
- Check for task status 
- Mark task as done 
- Update task information







## Tech

letsdo uses :

- Golang - Language 
- cobra - cli framework
- [MySQL] - Database

## Installation

letsdo requires [Golang](https://golang.org/),[Cobra, and MySQL](https://github.com/spf13/cobra).

Install the GO and run.

```sh
go get https://github.com/pratikyadav33/CLI-letsdo
```
Make Sure your Mysql is up
## Commands

```sh
letsdo : will give you list of commands to use
letsdo add <Task-Name>   :   Adds Task 
letsdo check <-u> <-c>   :   Can be used to check for completed and ot completed task
letsdo done <Taks-id>    :  Mark your task as Done using the task Id displayed
letsdo update <-a>   :   Delete all your Task
letsdo update <id> <-d>  :   Delete a Specifc Task
letsdo update <id> <Task-Name> <-t>   : Edit Your Task Name
letsdo update <Task Id> <1> <0>  -s  :  Change the status of your task
```
