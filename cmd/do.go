package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
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
		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatal(err)
		}

		//   the keys of the tasks to be removed
		for _, finishedTask := range finishedTasks {
			if finishedTask <= 0 || finishedTask > len(tasks) {
				fmt.Println("Invalid task number:", finishedTask)
				continue
			}
			task := tasks[finishedTask-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", finishedTask)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
