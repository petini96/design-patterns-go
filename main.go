package main

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {

	fmt.Print("first commit with session structure!")
}
