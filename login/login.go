package login

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


/**
 * LoginCall checks the user credentials against database
    if the credentials exist the function will return
    the user structure with specidied ID and username
 */
func LoginCall(user userlib.User) (userlib.User, bool) {
        // declare the database variable
        var database *sql.DB

        // check that the database exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // the database does not exist, create a new one
                userlib.CreateDataBase(database)
                fmt.Println("create new user with 'newuser' command\n")
                return user, false
        }

        // obtain login credentials
        username, password := getCreds()
        filename := username + "_log.txt"

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
                return user, false
        case nil:
                return userlib.User{id, username, filename}, true
        default:
                log.Fatal(err)
        }

        return user, false
}


/**
 * NewUser generates a new profile and insert user credentials into the database
 */
func NewUser(user userlib.User) userlib.User {
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
        newuser := insertTable(database, username, password)
        database.Close()

        return newuser
}


/**
 * insertTable inserts the given username and password to the database table
 */
func insertTable(db *sql.DB, username string, password string) userlib.User {
        fmt.Println("generating new user profile...")
        filename := username + "_log.txt"

        insert := `INSERT INTO users(Username, Password, Filename) VALUES (?, ?, ?)`
        // prepare to insert information into the table of the database
        statement, err := db.Prepare(insert)
        errCheck(err)
        _, err = statement.Exec(username, password, filename)
        errCheck(err)

        // select id from database
        row := db.QueryRow("SELECT id FROM users WHERE Username=$1 AND Password=$2",
                username, password)
        var id int
        switch err := row.Scan(&id); err {
        case sql.ErrNoRows:
                fmt.Println("Could not create user...")
                return userlib.User{0, "guest", ""}
        case nil:
                fmt.Println("New user has been created!\n")
                return userlib.User{id, username, filename}
        default:
                log.Fatal(err)
        }
        return userlib.User{0, "guest", ""}
}


/**
 * getCreds reads in credentials from user and returns them as strings
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


/**
 * DisplayUser displays the current logged in user
 */
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
