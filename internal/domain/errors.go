package domain

import "fmt"

var (
	PersonNotFound     = fmt.Errorf("not found")
	MovieNotFound      = fmt.Errorf("not found")
	ProfessionNotFound = fmt.Errorf("not found")
)
