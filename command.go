package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

func addTask(description string) (string, error) {
	tasks, err := loadTasks()
	if err != nil {
		return "", err
	}

	var id int

	if len(tasks) == 0 {
		id = 1
	} else {
		size := len(tasks) - 1
		id = tasks[size].Id + 1
	}

	task := NewTask(description)
	task.Id = id

	tasks = append(tasks, task)

	err = saveTasks(tasks)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task added successfully (ID : %d)",task.Id),nil
}

func listTask()(string,error){
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}

	var result string 

	if len(tasks) == 0{
		return "NO TASK FOUND",nil
	}else{
		for _,t:= range tasks{
			result += fmt.Sprintf("%d %s [%s]\n",t.Id,t.Description,t.Status)
		}
	}
	return result,nil
}

func updateTask(id int ,des string) (string,error) {
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}

	found := false
	for i := range tasks{
		if tasks[i].Id == id{
			tasks[i].Description = des
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found{
		return "Task Not Found",nil
	}

	err = saveTasks(tasks)
	if err != nil{
		return "",err
	}else{
		return "Task updated successfully",nil
	}
}

func deleteTask(id int)(string,error){
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}

	found := false

	for i:= range tasks{
		if tasks[i].Id == id {
			tasks = slices.Delete(tasks,i,i+1)
			found = true 
			break
		}
	}
	if !found{
		return "Task Not Found",nil
	}
		err = saveTasks(tasks)
			if err != nil{
				return "",err
			}
		return "Task deleted Successfully",nil
}

func MarkDone(id int)(string,error){
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}

	found := false

	for i:= range tasks{
		if tasks[i].Id == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
			found = true 
			break
		}
	}
	if !found{
		return "Task Not Found",nil
	}

	err = saveTasks(tasks)
			if err != nil{
				return "",err
			}
		return "Task Marked Successfully",nil
}

func MarkInProgress(id int)(string,error){
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}

	found := false

	for i:= range tasks{
		if tasks[i].Id == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now()
			found = true 
			break
		}
	}
	if !found{
		return "Task Not Found",nil
	}

	err = saveTasks(tasks)
			if err != nil{
				return "",err
			}
		return "Task Marked Successfully",nil

}

func listTaskByStatus(status string) (string,error){
	tasks,err := loadTasks()
	if err != nil{
		return "",err
	}
	var result string

	found := false

	for i := range tasks{
		if strings.EqualFold(tasks[i].Status,status){
			found = true
			result +=fmt.Sprintf("%d %s [%s]\n",tasks[i].Id,tasks[i].Description,tasks[i].Status)
		}
	}
	if !found{
		return "No Task found",nil
	}
	return result,nil

}

