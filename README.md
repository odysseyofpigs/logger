# LoggerApp

Simple logger application that acts as a simplified logging/blogging software.
Information is read and stored to a target file.

So far the program allows the user to log into the system if they have a 
profile, or create one if they desire. All user information is stored within a 
database file locally. The database is read using sqlite3 and so require the 
sqlite3 drivers listed under dependencies. The program Creates [user]_log.txt 
files for each individual user within the system so that writes can be stored 
independently.

The program allows you to create a new log under your user profile and store it 
within a personal [user]_log.db file within the generated 'logs' directory. The program 
will read up to a set buffer size and append it to the database file under the specified 
username, date, and time. The program currently does not support log export functionality 
and can only be read from within the program.


# Dependencies

Golang version 1.14: https://golang.org/dl/

https://github.com/mattn/go-sqlite3

# Installation

```bash
go get github.com/odysseyofpigs/loggerapplication
```

# TODO

* I need to create comments at the top of the files that dictate the functionality of the package.
* Create function that reads out information from log.db files and writes them to 
easily readable text files for storage or some other operation.
* Create a soft link file for the application. This requires rewrite of database reading 
and writing funcitons within the program.
