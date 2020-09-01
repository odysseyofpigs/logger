package userinfo

import (
        "fmt"
        "os/user"
)


// FindUserName obtains the username from the current computer session and
// returns it as a string
func FindUserName() string {
        user, err := user.Current()
        if err != nil {
                fmt.Println("Error: unable to obtain username")
                return ""
        }
        return user.Username
}
