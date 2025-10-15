/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes task, it will not be marked as completed.",
	Run: func(cmd *cobra.Command, args []string) {
		// Converting input by user to int and creating list of task to be removed
		var finishedTasks []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				finishedTasks = append(finishedTasks, id)
			}
		}

		// Looping to all list task to remove the ones marked by the user
		tasks, err := db.AllTasks(db.TaskBucket)
		if err != nil {
			log.Fatal(err)
		}

		// Looping through tasks marked as completed
		for _, finishedTask := range finishedTasks {
			if finishedTask <= 0 || finishedTask > len(tasks) {
				fmt.Println("Invalid task number:", finishedTask)
				continue
			}
			task := tasks[finishedTask-1]
			// Deleting task from Pending task bucket
			err = db.DeleteTask(task.Key)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("You have deleted the \"%s\" task.\n", task.Value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
