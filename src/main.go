package main

import (
        "fmt"
        "os"
    )

func main() {

    args := os.Args[1:]
    configPath, err := getInitPath()
    if err != nil{
        return
    }

    if _, err := os.Stat(configPath); args[0] != "init" && os.IsNotExist(err) {
        fmt.Println("file does not exist")
        os.Exit(1)
    }
    todoList, err := loadInit()
    if err != nil{
        return
    }

    switch command := args[0]; command {
        case "init":
            makeInit()
        case "a":
            todoList.addEntry(args[1:])
        case "l":
            todoList.showEntries(args[1:])
        case "c":
            todoList.checkEntry(args[1:])
        case "d":
            todoList.deleteEntry(args[1:])
        case "u":
            todoList.uncheckEntry(args[1:])
        }

    /*
    if (args[0] == "init") {
        makeInit()
    } else if (args[0] == "a") {
        todoList.addEntry(args[1:])
    } else if (args[0] == "l") {
        todoList.showEntries(args[1:])
    } else if (args[0] == "c") {
        todoList.checkEntry(args[1:])
    } else if (args[0] == "d") {
        todoList.deleteEntry(args[1:])
    } else if (args[0] == "u") {
        todoList.uncheckEntry(args[1:])
    }*/

}
