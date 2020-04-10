package login

/**
 * @file login.go
 * @author odysseyofpigs
 * @description This file contains all login functionality for the current User
 *  session.
 * @functionality The library handles the login function for the User. The library
 *  takes stdin from the terminal in the form of username and password credentials
 *  to determine whether the user can log into their respective session.
 */


import (
        "fmt"
        "log"
        "strings"
        "os"
        "bufio"
        "database/sql"
        "github.com/odysseyofpigs/loggerapplication/userlib"
        // driver for sqlite3 operations
        _ "github.com/mattn/go-sqlite3"
)



// LoginCall checks the user credentials against database
// if the credentials exist the function will return
// true and change the user struct credentials, false otherwise
func LoginCall(user *userlib.User) bool {
        // declare the database variable
        var database *sql.DB

        // check that the database exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // the database does not exist, create a new one
                userlib.CreateDataBase(database)
                fmt.Println("create new user with 'newuser' command\n")
                return false
        }

        // obtain login credentials
        username, password := getCreds()
        filename := username + "_log.db"

        // create a connection with the database
        database, _ = sql.Open("sqlite3", "./userlog.db")
        defer database.Close()

        // Query through table to see if login credentials match
        st := `SELECT id FROM users WHERE Username=$1 AND Password=$2`
        row := database.QueryRow(st, username, password)
        var id int
        switch err := row.Scan(&id); err {
        case sql.ErrNoRows:
                fmt.Println("user does not exist")
                fmt.Println("create user with 'newuser' command\n")
                return false
        case nil:
                user.ID = id
                user.Username = username
                user.Filename = filename
                return true
        default:
                log.Fatal(err)
        }

        return false
}



// NewUser generates a new profile and insert user credentials into the database
// the new user is automatically logged in to the system
func NewUser(user *userlib.User) {
        // declare the database variable
        var database *sql.DB

        // check that the database exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // if userlog does not exist, create database
                userlib.CreateDataBase(database)
        }

        // connect to the userlog database
        database, _ = sql.Open("sqlite3", "./userlog.db")

        // insert new user to database
        username, password := getCreds()
        insertTable(database, username, password)
        // select id from database
        row := database.QueryRow("SELECT id FROM users WHERE Username=$1 AND Password=$2",
                username, password)
        var id int
        switch err := row.Scan(&id); err {
        case sql.ErrNoRows:
                fmt.Println("User generation...failure")
                fmt.Println("Could not find new user...exiting")
                os.Exit(1)
        case nil:
                fmt.Println("New user has been created!\n")
                user.ID = id
                user.Username = username
                user.Filename = username + "_log.db"
        default:
                log.Fatal(err)
        }
        database.Close()
}


/**
 * insertTable inserts the given username and password to the database table
 */
func insertTable(db *sql.DB, username string, password string) {
        fmt.Println("generating new user profile...")
        filename := username + "_log.db"

        insert := `INSERT INTO users(Username, Password, Filename) VALUES (?, ?, ?)`
        // prepare to insert information into the table of the database
        statement, err := db.Prepare(insert)
        errCheck(err)
        _, err = statement.Exec(username, password, filename)
        errCheck(err)
}


/**
 * getCreds reads in credentials from stdin and returns them as strings
 */
func getCreds() (string, string) {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("username: ")
        u, err1 := reader.ReadString('\n')
        errCheck(err1)
        fmt.Print("password: ")
        p, err2 := reader.ReadString('\n')
        errCheck(err2)
        username := strings.Trim(u, "\n")
        password := strings.Trim(p, "\n")
        return username, password
}



// DisplayUser displays the current logged in user
func DisplayUser(user userlib.User) {
        fmt.Println("User Logged in::")
        fmt.Println("--------------------")
        fmt.Printf("ID      : %d\n", user.ID)
        fmt.Printf("Username: %s\n", user.Username)
        fmt.Printf("log file: %s\n", user.Filename)
        fmt.Print("\n")
}


// errCheck checks if any errors have occured
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
