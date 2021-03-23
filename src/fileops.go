package main


import (
        "strconv"
        "fmt"
        "encoding/json"
        "os"
    )

type Task struct{
    TaskID          int     `json:"taskID"`
    TaskDescription string  `json:"taskDescription"`
    Project         string  `json:"project"`
    TaskStatus      bool    `json:"status"`
    Due             string  `json:"Due"`
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


func render(task *Task) {

    var check string
    if task.TaskStatus {
        check = "[x]"
    } else {
        check = "[ ]"
    }
    fmt.Printf("%d      %s   %s\n", task.TaskID, check, task.TaskDescription)

}

func (todoList *Todo) addEntry(args []string) {
    description := args[0]
    newTask := &Task{}

    todoList.CurrentID = todoList.CurrentID + 1
    newTask.TaskID = todoList.CurrentID
    newTask.TaskDescription = description
    newTask.TaskStatus = false

    todoList.Tasks = append(todoList.Tasks, newTask)
    todoList.saveInit()

}


func (todoList *Todo) showEntries() {

    for _, val := range todoList.Tasks {
        render(val)
    }


}

func (todoList *Todo) checkEntry(args []string) {

    taskNum, _ := strconv.Atoi(args[0])
    for _, task := range todoList.Tasks {

        if task.TaskID == taskNum {
            task.TaskStatus = true
            continue
        }
    }
    todoList.saveInit()
}

func (todoList *Todo) deleteEntry(args []string) {

    var taskIdx int
    taskNum, _ := strconv.Atoi(args[0])

    for idx, task := range todoList.Tasks{
        if task.TaskID == taskNum{
            taskIdx = idx
        }
        if task.TaskID > taskNum{
            task.TaskID = task.TaskID - 1
        }
    }
    todoList.CurrentID = todoList.CurrentID - 1
    todoList.Tasks = append(todoList.Tasks[:taskIdx], todoList.Tasks[taskIdx + 1:]...)
    todoList.saveInit()
}

