package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RusimbiPatrick/contact-management-oreilly/contact"
)

func displayHelp(){
    fmt.Println("Available commands:")
    fmt.Println("1. add <name> <phone number> <email> - Add a new contact.")
    fmt.Println("2. view <name>   - View a contact's details.")
    fmt.Println("3. delete <name> - Delete a contact.")
    fmt.Println("4. list - List all contacts.")
    fmt.Println("5. help - Display this help message.")
    fmt.Println("6. exit - Exit the program.")
    fmt.Println("Please enter a command:")
}
func main(){
	fmt.Println("Welcome to the Contact Management System")
	displayHelp()
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("\n Enter command: ")
		scanner.Scan()
		command := scanner.Text()
		args := strings.Fields(command)
		if len(args) == 0 {
			fmt.Println("No command entered. Please try again.")
			continue
		}

		switch args[0] {
		case "add":
			if len(args) != 4 {
				fmt.Println("Usage: add <name> <phone number> <email>")
				continue
			}
			name := args[1]
			phoneNumber := args[2]
			email := args[3]
			contact := contact.Contact{Name: name, PhoneNumber: phoneNumber, Email: email}
			err := contact.AddContact(contact)
			if err != nil {
				fmt.Println("Error adding contact:", err)
			} else {
				fmt.Println("Contact added successfully.")
			}
		case "view":
			if len(args) != 2 {
				fmt.Println("Usage: view <name>")
				continue
			}
			name := args[1]
			contact, err := contact.ViewContact(name)
			if err != nil {
				fmt.Println("Error viewing contact:", err)
			} else {
				fmt.Printf("Contact Details - Name: %s, Phone Number: %s, Email: %s\n", contact.Name, contact.PhoneNumber, contact.Email)
			}
		case "delete":
			if len(args) != 2 {	
				fmt.Println("Usage: delete <name>")
				continue
			}
			name := args[1]
			err := contact.DeleteContact(name)
			if err != nil {
				fmt.Println("Error deleting contact:", err)
			} else {
				fmt.Println("Contact deleted successfully.")
			}
		case "list":
			contacts := contact.GetAllContacts()
			if len(contacts) == 0 {
				fmt.Println("No contacts found.")
			} else {
				fmt.Println("Contacts:")
				for _, c := range contacts {
					fmt.Printf("Name: %s, Phone Number: %s, Email: %s\n", c.Name, c.PhoneNumber, c.Email)
				}
			}
		case "help":
			displayHelp()
		case "exit":
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Unknown command:", args[0])
			fmt.Println("Type 'help' for a list of available commands.")		
		}
	}
}