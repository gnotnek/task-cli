package cmd

import "github.com/spf13/cobra"

func Execute() {
	var command = cobra.Command{
		Use:   "taskQ",
		Short: "Task manager",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(addTaskCmd(), getTaskByIDCmd(), getTasksCmd(), updateTask(), deleteTask())

	if err := command.Execute(); err != nil {
		panic(err)
	}
}
