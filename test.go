package main

import (
        "fmt"
        s  "strings"
)

var p = fmt.Println

func main() {
	ou := "x509::CN=User1@org1.example.com,OU=client,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US"
	//user := make([]string, 12) 
    user := s.Split(ou, ",")
	fmt.Println(user[0])	
	test := user[0]
	getuser := test[9:]
	fmt.Println(getuser)
}