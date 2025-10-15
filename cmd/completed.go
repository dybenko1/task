/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"task/db"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows today's completed tasks",

	Run: func(cmd *cobra.Command, args []string) {
		CompletedTasks, err := db.AllTasks(db.CompletedBucket)
		if err != nil {
			log.Fatal(err)
		}

		DateTasks, err := db.AllTasks(db.DateCompletion)
		if err != nil {
			log.Fatal(err)
		}

		// Identifying teh tasks completed today:
		// From the fact that the completed and date nuckets have the same index we can do this, if not we would need to first retrieve the key of elements tht have todays date from the date bucket and then search those keys in the completed bucket
		var todayCompletedTasks []db.Task
		for i, completedTask := range DateTasks {
			if completedTask.Value == db.TodaysDate() {
				task := CompletedTasks[i]
				todayCompletedTasks = append(todayCompletedTasks, task)
			}
		}
		if len(todayCompletedTasks) == 0 {
			fmt.Println("You have not completed any task to day. Tlabaja!!")
		} else {
			fmt.Println("You have finished the following tasks today:")
			for _, task := range todayCompletedTasks {
				fmt.Printf("- %s\n", task.Value)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
