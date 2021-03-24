package main

import (
        "os"
        "encoding/json"
        "os/user"
        "path"
    )


func checkInit() bool {
    return false
}

func getInitPath() (string, error) {
    confPath := "/.config/godo/godo.json"
    usr, err := user.Current()
    if err != nil {
        return "Error in obtaining current user", err
    }
    return path.Join(usr.HomeDir, confPath), nil
}
func loadInit() (Todo, error) {
    confPath := "/.config/godo/godo.json"
    usr, err := user.Current()
    if err != nil{
        return Todo{}, err
    }

    initPath := path.Join(usr.HomeDir, confPath)
    initFile, _ := os.ReadFile(initPath)
    var todo Todo
    json.Unmarshal(initFile, &todo)
    return todo, nil
}

func makeInit() {
    confPath := "/.config/godo/godo.json"
    usr, err := user.Current()
    if err != nil {
        return
    }
    initPath := path.Join(usr.HomeDir, confPath)
    data := Todo{SavePath: initPath, Tasks: make([]*Task,0)}
    file, _ := json.MarshalIndent(data, "", "")
    _ = os.WriteFile(initPath, file, 0644)
}

func (todoList *Todo) saveInit() {
    file, _ := json.MarshalIndent(todoList, "", "")
    _ = os.WriteFile(todoList.SavePath, file, 0644)
}




