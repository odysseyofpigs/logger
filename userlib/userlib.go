package userlib

import (
        "os"
        "log"
        "fmt"
        "database/sql"
        // driver for sqlite3 operations
        _ "github.com/mattn/go-sqlite3"
)

// User contains the current client information
type User struct {
        ID int
        Username string
        Filename string
}


// ListAll lists all the users within the database
func ListAll() {
        var database *sql.DB

        // check if userlog.db exists
        if _, err := os.Stat("userlog.db"); os.IsNotExist(err) {
                // create the database and return
                CreateDataBase(database)
                fmt.Println("Error: no users within database\n")
        } else {
                fmt.Println("Users::")
                fmt.Println("----------------")
                // print all the users within the database
                database, _ = sql.Open("sqlite3", "./userlog.db")
                defer database.Close()

                rows, err := database.Query("SELECT id, Username FROM users")
                errCheck(err)
                var id int
                var username string
                for rows.Next() {
                        // store the information within the variables listed
                        rows.Scan(&id, &username)
                        fmt.Printf("ID: %d | Username: %s\n", id, username)
                }
                fmt.Print("\n")
        }
}


// CreateDataBase creates a new database file and calls to initialize a
// completely new table
func CreateDataBase(db *sql.DB) {
        fmt.Println("No database found....creating new database file")
        // creat the new database file
        file, err := os.Create("userlog.db")
        errCheck(err)
        fmt.Println("database creation...complete")
        file.Close()

        // initialize the database with a new table
        db, _ = sql.Open("sqlite3", "./userlog.db")
        defer db.Close()

        // call table creation
        createTable(db)
        fmt.Println("database table initialization...complete")
}


// createTable initializes the given database with a new table
func createTable(db *sql.DB) {
        create := `CREATE TABLE users (
                "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                "Username" TEXT,
                "Password" TEXT,
                "Filename" TEXT
                );`

        statement, err := db.Prepare(create)
        errCheck(err)
        // execute the prepared table statement
        statement.Exec()
}


// errCheck checks if any errors have occured
func errCheck(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
