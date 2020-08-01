package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Users struct{
	Cids []string `json:"cids"`
}

func main() {
	fmt.Println("Starting to fetch users")
	//users, err := GetITUsers()
	users, err := ExtractUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", users.Cids)
	PostUsers(users)
}

func PostUsers(users Users) {
	client := &http.Client{}
	jsonUsers, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/users/whitelist", config.Address),
		bytes.NewBuffer(jsonUsers))
	if err != nil{
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf( "pre-shared %s", config.ApiKey))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	print(req.GetBody)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("completed request")
	fmt.Print(resp.Status)
}