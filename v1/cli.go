package v1

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func RunCLI(store *ToDoStore) {
	var (
		add    = flag.NewFlagSet("add", flag.ExitOnError)
		remove = flag.NewFlagSet("remove", flag.ExitOnError)
		list   = flag.NewFlagSet("list", flag.ExitOnError)
	)

	addTitle := add.String("title", "", "Title of the ToDo")
	addDue := add.String("due", "", "Due datetime (optional, format: 2006-01-02T15:04:05)")

	removeID := remove.Int("id", 0, "ID of the item to remove")

	listStatus := list.String("status", "", "Filter by status (NotStarted, InProgress, Completed)")
	listScope := list.String("scope", "", "Scope: archive or overdue")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'remove' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])

		var dueTime *time.Time
		if *addDue != "" {
			t, err := time.Parse("2006-01-02T15:04:05", *addDue)
			if err != nil {
				fmt.Println("Invalid datetime format. Use: 2006-01-02T15:04:05")
				os.Exit(1)
			}
			dueTime = &t
		}

		id := store.Add(*addTitle, dueTime)
		fmt.Printf("Added ToDo item with ID %d\n", id)

	case "remove":
		remove.Parse(os.Args[2:])
		if *removeID == 0 {
			fmt.Println("Please provide a valid --id to remove")
			os.Exit(1)
		}
		store.Remove(*removeID)
		fmt.Printf("Removed ToDo item with ID %d\n", *removeID)

	case "list":
		list.Parse(os.Args[2:])
		items := store.List(*listStatus, *listScope)
		for _, item := range items {
			fmt.Printf("ID: %d | Title: %s | Status: %s | Due: %v\n", item.ID, item.Title, item.Status, item.Due)
		}

	default:
		fmt.Println("expected 'add', 'remove' or 'list' subcommands")
		os.Exit(1)
	}
}
