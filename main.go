package main

import (
	"log"

	"github.com/spf13/cobra"
)


func main() {
	log.SetFlags(log.LUTC)

	
	var rootCmd = &cobra.Command{Use: "taskManager"}

	var createTask = &cobra.Command{
		Use: "create",
		Short: "Create a new task",
		Run: func (cmd *cobra.Command, args []string) {
			create()
			
		},
	
	}
	
	
	var readTask = &cobra.Command{
		Use: "read",
		Short: "Show the current tasks",
		Run: func(cmd *cobra.Command, args []string) {
			read()
			
		},
	
	}


	var insertTask = &cobra.Command{
		Use: "insert",
		Short: "Add a new task",
		Run: func(cmd *cobra.Command, args []string) {
			insert()

		},

	}	


	var updateTask = &cobra.Command{
		Use: "update",
		Short: "Update a task",
		Run: func(cmd *cobra.Command, args []string){
			update()

		},

	}


	var deleteTask = &cobra.Command{
		Use: "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string){
			delete()

		},

	}
	

	rootCmd.AddCommand(createTask, readTask, insertTask, deleteTask, updateTask)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)

	}
}