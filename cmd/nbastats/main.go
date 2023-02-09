package main

import (
	"fmt"
	"nbastats/internal/pkg/datamodel"
)

func main() {
	team, _ := datamodel.GetTeam(1610612737)

	fmt.Print(*team)
}
