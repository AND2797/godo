package main


import (
        "strconv"
        "fmt"
    )

type Task struct{
    TaskID          int     `json:"taskID"`
    TaskDescription string  `json:"taskDescription"`
    TaskProject     string  `json:"project"`
    TaskStatus      bool    `json:"status"`
    Due             string  `json:"Due"`
}


type Todo struct{
    SavePath      string   `json:"config path"`
    CurrentID     int      `json:"currentID"`
    Tasks         []*Task  `json:"Task"`
}



func render(task *Task) {
    var check string
    var project string
    if task.TaskStatus {
        check = "[x]"
    } else {
        check = "[ ]"
    }

    if task.TaskProject != ""{
        project = fmt.Sprintf("project: %s", task.TaskProject)
    }
    fmt.Printf("%d      %s   %s     %s\n", task.TaskID, check, task.TaskDescription, project)
}

func (todoList *Todo) addEntry(args []string) {
    description := args[0]
    newTask := &Task{}
    if len(args) >= 2 && args[1] == "p" {
        newTask.TaskProject = args[2][2:]
    }

    todoList.CurrentID = todoList.CurrentID + 1
    newTask.TaskID = todoList.CurrentID
    newTask.TaskDescription = description
    newTask.TaskStatus = false
    todoList.Tasks = append(todoList.Tasks, newTask)
    todoList.saveInit()
}

func (todoList *Todo) showProjects(project string){
    for _, val := range todoList.Tasks {
        if val.TaskProject == project{
            render(val)
        }
    }
}

func (todoList *Todo) showEntries(args []string) {
    argLength := len(args)

    switch argLength := argLength; argLength{
        case 0:
            if len(args) == 0{
                for _, val := range todoList.Tasks {
                    render(val)
                }
            }
        case 1:
            cueArgs := args[0]
            switch cuePrompt := args[0][0:2]; cuePrompt{
                case "p:":
                    todoList.showProjects(cueArgs[2:])
                default:
                    fmt.Println("no case match.")
                }
        default:
            fmt.Println("default")

    }
}

func (todoList *Todo) checkEntry(args []string) {
    taskNum, _ := strconv.Atoi(args[0])
    for _, task := range todoList.Tasks {

        if task.TaskID == taskNum {
            task.TaskStatus = true
            break
        }
    }
    todoList.saveInit()
}

func (todoList *Todo) uncheckEntry(args []string) {

    taskNum, _ := strconv.Atoi(args[0])
    for _, task := range todoList.Tasks {

        if task.TaskID == taskNum {
            task.TaskStatus = false
            break
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

