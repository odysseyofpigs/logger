package writelog

import (
        "fmt"
        "time"
        "os"
        "strings"
        "log"
        "bufio"
        "github.com/odysseyofpigs/loggerapplication/userlib"
)

/**
 * NewEntry creates a new entry to the log.txt file
 */
func NewEntry(user userlib.User) {
        // define file path
        prevDir, _ := os.Getwd()

        //check that the logs directory exists
        if _, err := os.Stat("logs"); os.IsNotExist(err) {
                // logs directory does not exist, create it
                createDir("logs")
        }

        //change the directory to the logs directory
        err := os.Chdir("logs")
        errCheck(err)

        //check to see if the [user]_log.txt file exists
        if _, err := os.Stat(user.Filename); os.IsNotExist(err) {
                // user log file does not exist
                fmt.Printf("%s file not found...creating new file\n", user.Filename)
                file, err := os.Create(user.Filename)
                errCheck(err)
                file.Close()
                fmt.Printf("%s created\n\n", user.Filename)
        }

        // create a new log within file
        writeToFile(user.Filename)

        // change the directory back to main directory
        err = os.Chdir(prevDir)
        errCheck(err)
}


/**
 * writeToFile takes input from standard input and writes it to the given file
 */
func writeToFile(filename string) {
        //open the file to append
        file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
        errCheck(err)
        defer file.Close()

        // get user name
        username := filename[:strings.IndexByte(filename, '_')]

        // get the time
        currentTime := time.Now()
        userEntry := "User " + username + " entry:: Time: " + currentTime.Format("2006.01.02 15:04:05") + "\n"

        // take user input
        fmt.Println("Enter log (limit to 4000 char):")
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">")
        logInput, _ := reader.ReadString('\n')

        a, err := file.Write([]byte(userEntry))
        errCheck(err)
        b, err1 := file.Write([]byte(logInput))
        errCheck(err1)

        fmt.Printf("Write complete with %d bytes\n\n", a+b)
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


// errCheck checks if any errors have occured
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
