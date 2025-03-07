package cmd

import (
	"taskQ/internal/task"

	"github.com/spf13/cobra"
)

func addTaskCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Run: func(cmd *cobra.Command, args []string) {
			task.AddTask(cmd.Flag("description").Value.String(), cmd.Flag("status").Value.String())
		},
	}

	command.Flags().StringP("description", "d", "", "Task description")
	command.Flags().StringP("status", "s", "pending", "Task status")
	return command

}

func getTaskByIDCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "get",
		Short: "Get a task",
		Run: func(cmd *cobra.Command, args []string) {
			task.GetTask(cmd.Flag("id").Value.String())
		},
	}

	command.Flags().StringP("id", "i", "", "Task ID")
	return command
}

func getTasksCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			task.ListTasks()
		},
	}

	return command
}

func updateTask() *cobra.Command {
	var command = &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Run: func(cmd *cobra.Command, args []string) {
			task.UpdateTask(cmd.Flag("id").Value.String(), cmd.Flag("description").Value.String(), cmd.Flag("status").Value.String())
		},
	}

	command.Flags().StringP("id", "i", "", "Task ID")
	command.Flags().StringP("description", "d", "", "Task description")
	command.Flags().StringP("status", "s", "", "Task status")
	return command
}

func deleteTask() *cobra.Command {
	var command = &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string) {
			task.DeleteTask(cmd.Flag("id").Value.String())
		},
	}

	command.Flags().StringP("id", "i", "", "Task ID")
	return command
}
