package main

import "os"

type Config struct {
	Address string
	ApiKey  string
}

var config = Config{
	Address: os.Getenv("ADDRESS"),
	ApiKey:  os.Getenv("API_KEY"),
}
