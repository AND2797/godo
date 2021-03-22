package main

import (
        "fmt"
        "log"
        "os"
    )


func main() {

    args := os.Args[1:]
    fmt.Println(args)

    if _, err := os.Stat("./test.json"); os.IsNotExist(err) {
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
        checkEntry(args[1:])
    } else if (args[0] == "d") {
        deleteEntry(args[1:])
    }

}
