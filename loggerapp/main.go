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
        "github.com/odysseyofpigs/loggerapplication/login"
        "github.com/odysseyofpigs/loggerapplication/writelog"
        "github.com/odysseyofpigs/loggerapplication/userlib"
)

func main() {
        fmt.Println("LoggerApp!")
        fmt.Println("Welcome, to the logger application")
        fmt.Println("type: 'help' for assistance\n")

        /* initialize user information */
        user := userlib.User{0, "guest", ""}

        /* read user input */
        input := readInput()

        /* check input until 'exit' is given */
        for input != "exit" {
                user = appCall(input, user)
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
func appCall(input string, user userlib.User) userlib.User {
        switch input {
        case "help":
                helpCall()
                break
        case "who":
                login.DisplayUser(user)
                break
        case "login":
                var loginStat bool
                loginStat = login.LoginCall(&user)
                if loginStat {
                        fmt.Printf("%s has logged in!\n\n", user.Username)
                }
                break
        case "newuser":
                login.NewUser(&user)
                break
        case "newlog":
                // check if the user is logged in
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        fmt.Printf("creating new log for %s\n", user.Username)
                        // create a new log entry for the user
                        writelog.NewEntry(user)
                }
                break
        case "logs":
                // check if the user is logged in
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        // call list log from userlib
                        fmt.Printf("Listing all log enteries for %s\n", user.Username)
                        writelog.ListLogs(user)
                }
                break
        case "export":
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        fmt.Println("Exporting log entires...")
                        writelog.ExportLogs(user)
                        fmt.Println("Log entries exported!\n")
                }
                break
        case "listall":
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        fmt.Print("\n")
                        userlib.ListAll()
                }
                break
        case "logout":
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        fmt.Printf("Logging %s out...\n", user.Username)
                        user.ID = 0
                        user.Username = "guest"
                        user.Filename = ""
                }
                break
        default:
                fmt.Println("::unknown command given::\n")
        } // end of switch
        return user
}


/*
 * helpCall the base help menu
 */
func helpCall() {
        fmt.Println("\nHelp Screen::")
        fmt.Println("--------------------------------------------")
        fmt.Println("login   :: login to your account")
        fmt.Println("newuser :: create a new user account")
        fmt.Println("who     :: lists current login session credentials")
        fmt.Println("\nLogged in Functionality::")
        fmt.Println("newlog  :: create a new log entry")
        fmt.Println("logs    :: list all log entries within the system")
        fmt.Println("export  :: exports all log enteries within system to txt file")
        fmt.Println("listall :: lists all users within the system")
        fmt.Println("logout  :: logs the current user out, changes to guest")
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
