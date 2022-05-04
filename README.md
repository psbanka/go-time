# State of our code:

- We have 4 branches. What do they do?
- main, gia-main, wild-west, 21Jul20-implement-basic-db-save
- We decided to structure the code based off of what we saw in [proteusdns app](https://github.com/fastly/proteusdns)


# Useful commands:

- `go fmt` does formatting stuffs
- `go build` builds `main.go` but you can also say `go build forms.go` and that will just build that thing.
  - also result in the creation of a binary executable by the same name

  - then `go build` will look seek out and find `func main()` in your current
  directory and it will also consult your `go.mod` file to figure out the name
  of the executable to create. If you have more than one `func main()` then `go
  build` will refuse to build.
  
  - `go run` builds without creating an executable file and just runs it (without creating the executable file)
  - `go help ___` will give you (possibly more than you want) help about a given command.


# Packages!

- We see that we have a `package main` at the top of some files. What is that about?

  A: if a go module is to be executable on its own, it needs to be in `package main`. if it's a library, it could be a different package.

- How do you import packages within the same repository?

- How would you make another module and import it? would it also have `package main`? at what point would it become `package kitties` does `package main` also need `func main()`?

- what is the work of `package` what is the work of `main`

- Does every project have to have one and only one `func main()` ?


TODOs: 
- Look at some kind of "real" projects:
  - look at IAM's thing as an API!
  - Go look at snit code!
  - We looked at proteusdns and immediately started fan-personing
- How does configuration within a library work? How do we do config/secrets management (e.g. root database passwords)

- How do we install dependencies?
- How does "github.com/go-sql-drive/mysql" get there? What did we have to do?
- How does package mgmt work in go?
- namespaces: when I import `fmt`, what did I get?
- what's the difference between
  - import "database/sql"
  - import "github.com/go-sql-driver/mysql"
  - The difference is that the first one is a standard library whereas the mysql driver is an external library.

# Uncategorized musings:

	- // Question: how do you define constants? can you have immutable constants?

  It's as simple as using `const`.

	- defer rows.Close() // how does defer generally work?

  When you defer inside a function, you are telling the program to close whatever you had open (file, connection, etc.) after you have returned from the scope of the function.

  Does go have concept of exceptions? It must have runtime exceptions right? For example, what happens if you divide by zero?

	// WHat is this? How does it get set? Why do we care?
	err = rows.Err()

  - Type definitions: can types be declared outside the file they're used?

- how do we run stuff/ compile, etc?

- How do we lint stuff?