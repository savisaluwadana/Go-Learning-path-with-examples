package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-zero-to-hero/projects/phase-02-contact-manager/contacts"
)

func main() {
	manager := contacts.NewManager()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Contact Manager CLI")
	printHelp()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		cmd := strings.ToLower(args[0])

		switch cmd {
		case "add":
			if len(args) != 4 {
				fmt.Println("usage: add <name> <email> <phone>")
				continue
			}
			contact, err := manager.Add(args[1], args[2], args[3])
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("added contact #%d\n", contact.ID)
		case "list":
			list := manager.List()
			if len(list) == 0 {
				fmt.Println("no contacts found")
				continue
			}

			fmt.Println("ID | Name | Email | Phone")
			for _, c := range list {
				fmt.Printf("%d | %s | %s | %s\n", c.ID, c.Name, c.Email, c.Phone)
			}
		case "update":
			if len(args) != 5 {
				fmt.Println("usage: update <id> <name> <email> <phone>")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("error: id must be an integer")
				continue
			}
			updated, err := manager.Update(id, args[2], args[3], args[4])
			if err != nil {
				if errors.Is(err, contacts.ErrNotFound) {
					fmt.Println("error: contact not found")
					continue
				}
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("updated contact #%d (%s)\n", updated.ID, updated.Name)
		case "delete":
			if len(args) != 2 {
				fmt.Println("usage: delete <id>")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("error: id must be an integer")
				continue
			}
			if err := manager.Delete(id); err != nil {
				if errors.Is(err, contacts.ErrNotFound) {
					fmt.Println("error: contact not found")
					continue
				}
				fmt.Println("error:", err)
				continue
			}
			fmt.Printf("deleted contact #%d\n", id)
		case "help":
			printHelp()
		case "exit", "quit":
			return
		default:
			fmt.Println("unknown command. type 'help' for available commands")
		}
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  add <name> <email> <phone>")
	fmt.Println("  list")
	fmt.Println("  update <id> <name> <email> <phone>")
	fmt.Println("  delete <id>")
	fmt.Println("  help")
	fmt.Println("  exit")
}
