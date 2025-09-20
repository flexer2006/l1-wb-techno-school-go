package one

import "fmt"

type Human struct {
	field int
}

func (h Human) PPP() string {
	return fmt.Sprintf("Huma field = %d", h.field)
}

type Action struct {
	Human
}
