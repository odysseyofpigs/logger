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
within a personaalized [user]_log.txt file within the 'logs' directory. The program 
will read up to a set buffer size and append it to the .txt file under the specified 
username, date, and time.


# Dependencies

https://github.com/mattn/go-sqlite3

# Installation

```bash
go get github.com/odysseyofpigs/loggerapplication
```

# TODO

* Change log.txt files to be sql database files to instead store the information.
* Then, create function that reads out information from log.db files and writes them to 
easily readable text files for storage or some other operation.
