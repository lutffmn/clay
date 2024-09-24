

# Clay CLI

## A lightweight cli library written in Go to help build CLI Application



This library aims to help developer to build CLI based apps in simple and clear way. Has feature to adds long and short style flags.



## How to install this package

The easiest way to install this package is:

* Import to your project directly
  
  ```go
  import ("github.com/lutffmn/clay")
  ```

* Open terminal and navigate to your project directory and run

* ```shell
  go mod tidy
  ```

Or you can install it using 

* ```shell
  go get -u github.com/lutffmn/clay
  ```



## Usage

```go
package main

import (
	"fmt"

	"github.com/lutffmn/clay"
)

func main() {
	// Construct new app
	app := clay.New("my app", "app description", "my app [OPTION]", "v1.0.0")

	// Register a new feature using short style flag
	app.Short("p", simpleHandler)

	// Register a new feature using long style flag
	app.Long("pong", notSoSimpleHandler)

	// Parsing the args
	app.Parse()
}

// dummy handler
func simpleHandler() {
	fmt.Println("Ping")
}

// dummy handler
func notSoSimpleHandler(msg string) {
	fmt.Printf("%s", msg)
}

```



# Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above. If you would like to submit a PR with a fix, **please reference** the issue you created!


