package login

import (
        "fmt"
        "log"
        "strings"
        "os"
        "bufio"
        "database/sql"
        // driver for sqlite3 operations
        _ "github.com/mattn/go-sqlite3"
)

// User structure holds current user information
type User struct {
        ID int
        Username string
        Filename string
}


/**
 * LoginCall checks the user credentials against database
    if the credentials exist the function will return
    the user structure with specidied ID and username
 */
func LoginCall(user User) (User, bool) {
        // declare the database variable
        var database *sql.DB

        // check that the database exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // the database does not exist, create a new one
                CreateDataBase(database)
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
                return User{id, username, filename}, true
        default:
                log.Fatal(err)
        }

        return user, false
}


/**
 * NewUser generates a new profile and insert user credentials into the database
 */
func NewUser(user User) User {
        // declare the database variable
        var database *sql.DB

        // check that the database exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // if userlog does not exist, create database
                CreateDataBase(database)
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
 * CreateDataBase creates a new database file and initializes the table for it
 */
func CreateDataBase(db *sql.DB) {
        fmt.Println("No database found....creating new database")
        // create new database file
        file, err := os.Create("userlog.db")
        errCheck(err)
        fmt.Println("database creation...complete")
        //close the created file
        file.Close()

        // initialize the database with a new table
        db, _ = sql.Open("sqlite3", "./userlog.db")
        defer db.Close()

        //populate database with table
        createTable(db)
        fmt.Println("database table initialization...complete")
}


/**
 * createTable initializes the given database with a new table
 */
func createTable(db *sql.DB) {
        create := `CREATE TABLE users (
                "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                "Username" TEXT,
                "Password" TEXT,
                "Filename" TEXT
                );`

        statement, err := db.Prepare(create)
        errCheck(err)
        // execute the statement prepared
        statement.Exec()
}


/**
 * insertTable inserts the given username and password to the database table
 */
func insertTable(db *sql.DB, username string, password string) User {
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
                return User{0, "guest", ""}
        case nil:
                fmt.Println("New user has been created!\n")
                return User{id, username, filename}
        default:
                log.Fatal(err)
        }
        return User{0, "guest", ""}
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
func DisplayUser(user User) {
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
