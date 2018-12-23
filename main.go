package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//flags
var (
	user string
)

func main() {
	// Parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username:	`, result.Login)
		fmt.Println(`Name:		`, result.Name)
		fmt.Println(`Email:		`, result.Email)
		fmt.Println(`Bio:		`, result.Bio)
		fmt.Println("")
	}

}

func init() {
	flag.StringVar(&user, "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
