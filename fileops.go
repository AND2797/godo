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

func loadInit()  {
    initFile, _ := os.ReadFile("test.json")
    var todo Todo
    json.Unmarshal(initFile, &todo)
    fmt.Println(todo)


}

func makeInit() {

    data := Todo{Tasks: make([]*Task,0)}
    file, _ := json.MarshalIndent(data, "", "")

    _ = os.WriteFile("test.json", file, 0644)

}

func render() {

}

func addEntry() {

}


func showEntries() {
}

func checkEntry() {
}

func deleteEntry() {
}

