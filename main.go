package main

import (
        "fmt"
        "log"
        "os"
    )


func main() {

    args := os.Args[1:]
    //fmt.Println(args)

    if _, err := os.Stat("./test.json"); args[0] != "init" && os.IsNotExist(err) {
        fmt.Println("file does not exist")
        log.Fatal("end")
    }
    todoList := loadInit()

    if (args[0] == "init") {
        makeInit()
    } else if (args[0] == "a") {
        todoList.addEntry(args[1:])
    } else if (args[0] == "l") {
        todoList.showEntries()
    } else if (args[0] == "c") {
        todoList.checkEntry(args[1:])
    } else if (args[0] == "d") {
        todoList.deleteEntry(args[1:])
    }

}
