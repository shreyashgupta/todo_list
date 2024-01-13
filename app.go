package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the todo list",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() == 0 {
						fmt.Println("Error: Please specifiy a task")
					} else {
						todoList, err := ReadTodoListFromFile("./todo.json")
						if err != nil {
							return err
						}
						todoList.AddTask(cCtx.Args().First())
						if err != nil {
							return err
						}
						fmt.Println("added task: ", cCtx.Args().First())
					}
					return nil
				},
			},
			{
				Name:    "viewc",
				Aliases: []string{"vc"},
				Usage:   "view completed tasks",
				Action: func(cCtx *cli.Context) error {
					todoList, err := ReadTodoListFromFile("./todo.json")
					if err != nil {
						return err
					}
					todoList.PrintCompletedTasks()
					return nil
				},
			},
			{
				Name:    "viewa",
				Aliases: []string{"va"},
				Usage:   "view active tasks",
				Action: func(cCtx *cli.Context) error {
					todoList, err := ReadTodoListFromFile("./todo.json")
					if err != nil {
						return err
					}
					todoList.PrintActiveTasks()
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "mark a task as completed",
				Action: func(cCtx *cli.Context) error {
					if cCtx.NArg() == 0 {
						fmt.Println("Error: Please specifiy a task id")
					} else {
						todoList, err := ReadTodoListFromFile("./todo.json")
						if err != nil {
							return err
						}
						id, err := strconv.Atoi(cCtx.Args().First())
						if err != nil {
							return err
						}
						err = todoList.MarkTaskCompleted(id)
						if err != nil {
							return err
						}
						fmt.Printf("Marked task with index %d completed", id)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
