package writelog

import (
        "fmt"
        "os"
        "log"
)

/**
 * NewEntry creates a new entry to the log.txt file
 */
func NewEntry(user User) {
        // define file path
        prevDir, _ := os.Getwd()
        path := "logs/" + user.Filename

        //check that the logs directory exists
        if _, err := os.Stat("logs"); os.IsNotExist(err) {
                // logs directory does not exist, create it
                createDir("logs")
        }

        //change the directory to the logs directory
        err = os.Chdir("logs")
        errCheck(err)

        //check to see if the [user]_log.txt file exists
        if _, err := os.Stat(user.Filename); os.IsNotExist(err) {
                // user log file does not exist
                fmt.Printf("%s file not found...creating new file\n", user.Filename)
        }

        // change the directory back to main directory
        err = os.Chdir(prevDir)
        errCheck(err)
}


// errCheck checks if any errors have occured
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}


/**
 * createDir creates a new directory of given name
 */
func createDir(dirname string) {
        fmt.Printf("no '%s' directory found....creating new\n", dirname)
        err := os.Mkdir("logs", 0700)
        errCheck(err)
        fmt.Println("new directory...created")
}
