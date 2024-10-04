//go mod init filename

package main //package used to group functions and data together (pkg.go.dev) main is entry point of the program

import (
	"fmt" //format text package (console)

	"rsc.io/quote" //importing multiple packages
)

func main() { //main function, entry point of the program
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

}
