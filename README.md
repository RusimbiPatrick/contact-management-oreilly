# Simple CLI-Based Contacts Management System in Go

## Functional requirements
- Add a contact with fields such as name, phone number, and email.
- View contact details by searching with the contact's name.
- Delete a contact by providing the contact's name.
## Non-functional Requirements
- Graceful handling of errors with informative error messages.
- User-friendly command-line interface

## Modules
- main: Serves as the entry point of the program.
- contact: Responsible for handling contact operations(add, view,delete)
- validation: Handles input validation for phone numbers and other fields

## Data Structures
- Contact struct:
    - Fields: name, phoneNumber, email.
- Custom types: PhoneNumber for handling phone number validation.

## Methods:
- AddContact(contact Contact) error
- ViewContact(name string)(Contact, error)
- DeleteContact(name string) error
## Error Handling
### Potential error scenarios:
- Invalid input format for phone number or email addresses.
- Attempting to add a duplicate contact
- Custom error types and informative error messages will be provided
## User Interaction
The user will interact with the program through the command-line interface(CLI).
Example commands:
  - add John Doe 123-45607890 john.doe@example.com
  - view John Doe
  - delete John Doe

The Contacts Management System project focuses on applying core Go concepts practically. It aims to provide users with a simple yet effective tool for managing their contacts through a command-line interface.

