package main


import (
        "fmt"
        "encoding/json"
        "os"
    )

type Task struct{
    TaskID          int     `json:"taskID"`
    TaskDescription string  `json:"taskDescription"`
    Project         string  `json:"project"`
}


type Todo struct{
    CurrentID     int      `json:"currentID"`
    Tasks         []*Task  `json:"Task"`
}

func checkInit() bool {

    return false
}

func loadInit() (Todo) {
    initFile, _ := os.ReadFile("test.json")
    var todo Todo
    json.Unmarshal(initFile, &todo)
    fmt.Println(todo)
    return todo


}

func makeInit() {

    data := Todo{Tasks: make([]*Task,0)}
    file, _ := json.MarshalIndent(data, "", "")

    _ = os.WriteFile("test.json", file, 0644)

}

func (todoList *Todo) saveInit() {

    file, _ := json.MarshalIndent(todoList, "", "")
    _ = os.WriteFile("test.json", file, 0644)



}


func render() {

}

func (todoList *Todo) addEntry() {

    newTask := &Task{}
    newTask.TaskID = todoList.CurrentID + 1
    newTask.TaskDescription = "test"
    todoList.Tasks = append(todoList.Tasks, newTask)

}


func showEntries() {
}

func checkEntry() {
}

func deleteEntry() {
}

