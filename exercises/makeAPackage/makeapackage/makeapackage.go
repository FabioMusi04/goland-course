package makeapackage

import (
	"errors"
	"fmt"
)

func MakeAPackage(name string) (string, error) {
	//function that returns a string or an error and takes a string as an argument
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	packageCreated := fmt.Sprintf("Hello, %s!", name)
	return packageCreated, nil
}
