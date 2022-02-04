package main

import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	// pointer ke host, username, password (biar gak diduplikasi datanya)
	fmt.Println(*host, *username, *password)

}