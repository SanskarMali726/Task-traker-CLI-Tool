package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	if len(os.Args) < 2{
		fmt.Println("Give Proper command")
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command{
		case "add":
			des := strings.Join(os.Args[2:]," ")
			s,err:=addTask(des)
				if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(s)
			return 
		case "update":
			id,err:= strconv.Atoi(os.Args[2])
			if err != nil{
				fmt.Println(err)
				return 
			}
			des := strings.Join(os.Args[3:]," ")

			s,err := updateTask(id,des)
			if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(s)
			return 
		
		case "delete":
			id,err:= strconv.Atoi(os.Args[2])
			if err != nil{
				fmt.Println(err)
				return 
			}
			s,err := deleteTask(id)
			if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(s)
			return 

		case "mark-in-progress":
			id,err:= strconv.Atoi(os.Args[2])
			if err != nil{
				fmt.Println(err)
				return 
			}
			s,err := MarkInProgress(id)
			if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(s)
			return

		case "mark-done":
			id,err:= strconv.Atoi(os.Args[2])
			if err != nil{
				fmt.Println(err)
				return 
			}
			s,err := MarkDone(id)
			if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(s)
			return
		case "list":
			if len(os.Args) == 2{
				s,err := listTask()
				if err != nil{
					fmt.Println(err)
					return
				}
				fmt.Println(s)
				return
			}
			s,err := listTaskByStatus(os.Args[2])
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(s)
			return 
		default:
    		fmt.Println("Unknown command. Available: add, update, delete, mark-in-progress, mark-done, list")	
	}

}