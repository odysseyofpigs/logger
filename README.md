# LoggerApp

Simple logger application that acts as a simplified logging/blogging software.
Information is read and stored to a target file.

So far the program allows the user to log into the system if they have a 
profile, or create one if they desire. All user information is stored within a 
database file locally. The database is read using sqlite3 and so require the 
sqlite3 drivers.

Next step is to create [user]_log.txt files for each individual user within the 
system so that writes can be stored independently.

# Dependencies

https://github.com/mattn/go-sqlite3

# Optimization

Change functions to take User struct pointers to avoid construction of new User 
structures and returns. 
