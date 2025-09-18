package main

import (
	"encoding/json"
	"io"
	"os"

)
const taskfile = "task.json"

func loadTasks() ([]Task, error) {
	file,err := os.OpenFile(taskfile,os.O_CREATE|os.O_RDONLY,0644)
	if err != nil{
		return nil,err
	}
	defer file.Close()

	data,err := io.ReadAll(file)
	if err != nil{
		return nil,err
	}

	if len(data) == 0 {
        _, err := file.Write([]byte("[]"))
        if err != nil {
            return nil, err
        }
        return []Task{}, nil
    }

	var t []Task

	err = json.Unmarshal(data,&t)
	if err != nil{
		return nil,err
	}
	
	return t, nil
}

func saveTasks(tasks []Task) error {

	data,err := json.MarshalIndent(tasks,"","  ")
	if err != nil{
		return err
	}

	err = os.WriteFile("task.json.temp",data,0644)
	if err != nil{
		return err
	}

	err = os.Rename("task.json.temp",taskfile)
	if err != nil{
		return err
	}

	return nil
}


