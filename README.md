# State of our code:

- We have 4 branches. What do they do?
- main, gia-main, wild-west, 21Jul20-implement-basic-db-save


# Useful commands:

# Packages!

- We see that we have a `package main` at the top of some files. What is that about?
- How would you make another module and import it? would it also have `package main`? at what point would it become `package kitties` does `package main` also need `func main()`?
- what is the work of `package` what is the work of `main`
- Does every project have to have one and only one `func main()` ?


TODOs: 
- Look at some kind of "real" projects:
  - look at IAM's thing as an API!
  - Go look at snit code!
- How does configuration within a library work? How do we do config/secrets management (e.g. root database passwords)

- How do we install dependencies?
- How does "github.com/go-sql-drive/mysql" get there? What did we have to do?
- How does package mgmt work in go?
- namespaces: when I import `fmt`, what did I get?
- what's the difference between
  - import "database/sql"
  - import "github.com/go-sql-driver/mysql"

# Uncategorized musings:

	- // Question: how do you define constants? can you have immutable constants?

	- // TODO: Weird smell: Could this exist outside the main function?
	- // Could this be declared in another module and imported?

	- defer rows.Close() // how does defer generally work?

	// WHat is this? How does it get set? Why do we care?
	err = rows.Err()

  - Type definitions: can types be declared outside the file they're used?

- how do we run stuff/ compile, etc?

- How do we lint stuff?