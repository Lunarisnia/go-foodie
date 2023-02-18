package types

import "fmt"

type FailedCreateData struct {
}

func (e *FailedCreateData) Error() string {
	return fmt.Sprint("data creation failed")
}
