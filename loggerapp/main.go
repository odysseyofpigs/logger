package main

import (
        "fmt"
        "log"
        "strings"
        "bufio"
        "os"
        "github.com/odysseyofpigs/loggerapplication/login"
)

func main() {
        fmt.Println("LoggerApp!")
        fmt.Println("Welcome, to the logger application")
        fmt.Println("type: 'help' for assistance\n")

        /* initialize user information */
        user := login.User{0, "guest", ""}

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
func appCall(input string, user login.User) login.User {
        switch input {
        case "help":
                helpCall()
                break
        case "who":
                login.DisplayUser(user)
                break
        case "login":
                var loginStat bool
                user, loginStat = login.LoginCall(user)
                if loginStat {
                        fmt.Printf("%s has logged in!\n\n", user.Username)
                }
                break
        case "newuser":
                user = login.NewUser(user)
                break
        case "newlog":
                // check if the user is logged in
                if user.Username == "guest" && user.Filename == "" {
                        fmt.Println("Error: not logged in\n")
                } else {
                        fmt.Printf("creating new log for %s\n", user.Username)
                }
                break
        default:
                fmt.Println("::unknown command given::\n")
        } // end of switch
        return user
}


/**
 * helpCall the base help menu
 */
func helpCall() {
        fmt.Println("Help Screen::")
        fmt.Println("--------------------------------------------")
        fmt.Println("login   :: login to your account")
        fmt.Println("newuser :: create a new user account")
        fmt.Println("\nLogged in Functionality::")
        fmt.Println("newlog  :: create a new log entry")
}


/**
 * errCheck terminates the program is any error occurs within the process
 */
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
