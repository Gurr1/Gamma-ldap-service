package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExtractUsers() (Users, error){
	file, err := os.Open("ITsek.csv")
	if err != nil {
		fmt.Println("error ", err)
		return Users{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error ", err)
	}
	var users []string

	for i := 1; i< len(record); i++ {
		users = append(users, record[i][0])
	}

	return Users{Cids: users}, nil
}
