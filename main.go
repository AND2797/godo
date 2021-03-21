package main

import (
        "fmt"
        "os"
    )


func main() {

    makeInit()
    todoList := loadInit()
    todoList.addEntry()
    todoList.saveInit()
    todoList = loadInit()
    fmt.Println("loaded",todoList.Tasks[0].TaskID, todoList.Tasks[0].TaskDescription)
    argsWithoutProg := os.Args[1:]
    fmt.Println(argsWithoutProg)

//    if !checkInit() {
//
//    } else {
//        todoList := loadInit()
//    }
//
//    if (argsWithoutProg[0] == "init") {
//        makeInit()
//    } else if (argsWithoutProg[0] == "a") {
//        todoList.addEntry()
//    } else if (argsWithoutProg[0] == "l") {
//        showEntries()
//    } else if (argsWithoutProg[0] == "c") {
//        checkEntry()
//    } else if (argsWithoutProg[0] == "d") {
//        deleteEntry()
//    }

}
