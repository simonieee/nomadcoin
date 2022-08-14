package main

import (
	"fmt"

	"github.com/simonieee/nomadcoin/person"
)

func main() {
	simon := person.Person{}
	simon.SetDetails("simon", 29)
	fmt.Println("Main 'simon'", simon)
}
