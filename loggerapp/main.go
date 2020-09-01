package main

/**
 * @file main.go
 * @author odysseyofpigs
 * @description this file is the main file for the LoggerApplication.
 *  Compile and execute this file to run the application. The following
 *  libraries must be present within local the go workpace for the app to
 *  function.
 * @functionality This process contains the used userlib.User structure
 *  that holds current user information. The process takes user input and
 *  parses it, calling the appropriate functions from the libraries.
 */


import (
        "fmt"
        "log"
        "strings"
        "bufio"
        "os"
        "github.com/odysseyofpigs/loggerapplication/writelog"
        "github.com/odysseyofpigs/loggerapplication/userinfo"
)


// userInfo stores all user information
type userT struct {
        username string
}


func main() {
        fmt.Println("LoggerApp!")
        fmt.Println("Welcome, to the logger application")
        fmt.Println("type: 'help' for assistance\n")

        // initialize user information
        var user = userT{userinfo.FindUserName()}

        /* read user input */
        input := readInput()

        /* check input until 'exit' is given */
        for input != "exit" {
                appCall(input, user)
                input = readInput()
        }

        fmt.Println("Goodbye!\n")
}


/* -- functions -- */

/**
 * readInput reads given input from the user and stores each argument into
 *  a string array
 * @return split list of given input
 */
func readInput() string {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("$~")

        /* read input */
        input, err := reader.ReadString('\n')
        errCheck(err)
        userInput := strings.Trim(input, "\n")
        //usr := strings.Split(input, " ")

        return userInput
}


/**
 * appCall handles the function call given by the user
 */
func appCall(input string, user userT) {
        switch input {
        case "help":
                helpCall()
                break
        case "who":
                fmt.Println(user.username, "\n")
                break
        case "newlog":
                fmt.Printf("creating new log for %s\n", user.username)
                // create a new log entry for the user
                writelog.NewEntry(user.username)
                break
        case "listlogs":
                // call list log function from writelog
                fmt.Printf("Listing all log enteries for %s\n", user.username)
                writelog.ListLogs(user.username + "_log.db")
                break
        case "export":
                fmt.Println("Exporting log entires...")
                writelog.ExportLogs(user.username)
                fmt.Println("Log entries exported!\n")
                break
        default:
                fmt.Println("::unknown command given::\n")
        } // end of switch
}


/*
 * helpCall the base help menu
 */
func helpCall() {
        fmt.Println("\nHelp Screen::")
        fmt.Println("--------------------------------------------")
        fmt.Println("who     :: lists current system session credentials")
        fmt.Println("newlog  :: create a new log entry")
        fmt.Println("listlogs:: list all log entries within the system")
        fmt.Println("export  :: exports all log enteries within system to txt file")
        fmt.Println("exit    :: exits from the logger application")
        fmt.Print("\n")
}


/*
 * errCheck terminates the program is any error occurs within the process
 */
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
