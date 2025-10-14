/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()

}
