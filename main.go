package main


import (
        "fmt"
        "os"
    )


func main() {

    argsWithoutProg := os.Args[1:]
    fmt.Println(argsWithoutProg)

    if !checkInit() {

    } else {
        loadInit()
    }

    if (argsWithoutProg[0] == "init") {
        makeInit()
    } else if (argsWithoutProg[0] == "a") {
        addEntry()
    } else if (argsWithoutProg[0] == "l") {
        showEntries()
    } else if (argsWithoutProg[0] == "c") {
        checkEntry()
    } else if (argsWithoutProg[0] == "d") {
        deleteEntry()
    }

}
