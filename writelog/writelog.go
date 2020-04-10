package writelog

/**
 * @file writelog.go
 * @author odysseyofpigs
 * @description This file contains all nessecary functionality for talking with
 *  the User logs.
 * @functionality The library contains functions that work with the creation
 *  and initialization of the User log database file. The library also handles
 *  the creation and insertion of new log entries to the User log database.
 *  The library also handles the listing of all User log entries to the terminal.
 */


import (
        "fmt"
        "time"
        "os"
        "strings"
        "log"
        "bufio"
        "database/sql"
        "github.com/odysseyofpigs/loggerapplication/userlib"
        // driver for sqlite3 operations
        _ "github.com/mattn/go-sqlite3"
)

// NewEntry creates a new entry to the log.db file
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

        //check to see if the [user]_log.db file exists
        _ = checkdb(user.Filename)

        // create a new log within the database
        writeToFile(user.Filename)

        // change the directory back to main directory
        err = os.Chdir(prevDir)
        errCheck(err)
}


/**
 * writeToFile takes input from standard input and stores it within the
   user specific database.
 */
func writeToFile(filename string) {
        // open the database
        database, _ := sql.Open("sqlite3", "./"+filename)
        defer database.Close()

        // prepare statement to place in log
        fmt.Println("[Enter log]: ")
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">")
        logInput, _  := reader.ReadString('\n')

        // prepare login information and time
        currentTime := time.Now()
        username := filename[:strings.IndexByte(filename, '_')]
        userEntry := username + " entry || Time: " + currentTime.Format("2006.01.02 15:04:05") + "\n"

        // prepare database for log information
        insert := `INSERT INTO logs(date, note) VALUES (?, ?)`
        statement, err := database.Prepare(insert)
        errCheck(err)
        _, err = statement.Exec(userEntry, logInput)
        errCheck(err)

        // new log created and entered successfully
        fmt.Println("Log entered\n")


        // -- leave for now --
        /*
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
        */
}


// checkdb checks if the given database exists, if not creates it and
// populates the new database with a respective table. returns a bool
// value: true if the file exists, false otherwise
func checkdb(filename string) bool {
        // check if the database exists
        if _, err := os.Stat(filename); os.IsNotExist(err) {
                // does not exist, create it
                fmt.Printf("%s file not found...creating new file\n", filename)
                file, err := os.Create(filename)
                errCheck(err)
                file.Close()
                fmt.Printf("%s created\n\n", filename)

                // populate the database with a new table
                database, _ := sql.Open("sqlite3", "./"+filename)
                defer database.Close()
                setTable(database)
                fmt.Println("database table initialization...complete")
                return false
        }
        return true
}


// setTable sets the table within the given database
func setTable(db *sql.DB) {
        create := `CREATE TABLE logs (
                "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                "date" TEXT,
                "note" TEXT
                );`

        statement, err := db.Prepare(create)
        errCheck(err)
        statement.Exec()
}


// ListLogs lists all log entries for the given user
func ListLogs(user userlib.User) {
        // get prev dir
        prevDir, _ := os.Getwd()

        // check if logs even exists
        if _, err := os.Stat("logs"); os.IsNotExist(err) {
                createDir("logs")
                fmt.Println("Error: no users have been initialized")
                fmt.Println("use log database does not exist\n")
        } else {
                // change directory to logs
                err := os.Chdir("logs")
                errCheck(err)

                // check that the user_log.db file exists
                if checkdb(user.Filename) {
                        // open the database to read
                        database, _ := sql.Open("sqlite3", "./"+user.Filename)
                        defer database.Close()

                        // set variables to contain information from database
                        var id int
                        var time string
                        var note string

                        row, err := database.Query("SELECT id, date, note FROM logs")
                        errCheck(err)
                        fmt.Println("[Notes]::")
                        for row.Next() {
                                row.Scan(&id, &time, &note)
                                // print the information
                                fmt.Printf("%d : %s\n", id, time)
                                fmt.Printf("%s\n\n", note)
                        }
                } else {
                        // there is no information within the user log database
                        fmt.Println("Error: no information within database\n")
                }

                // list all logs within file

                // change the directory back
                err = os.Chdir(prevDir)
        }
}


/**
 * createDir creates a new directory of given name
 */
func createDir(dirname string) {
        fmt.Printf("no '%s' directory found....creating new directory\n", dirname)
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
