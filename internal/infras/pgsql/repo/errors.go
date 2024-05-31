package repo

import "fmt"

var CreationError = func(msg string) error {
	return fmt.Errorf("cannot create %s", msg)
}

var GetError = func(msg string) error {
	return fmt.Errorf("cannot get %s", msg)
}
