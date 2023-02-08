package main

import (
	"context"
	"fmt"
	"nbastats/internal/pkg/datamodel"
)

func main() {
	ctx := context.Background()
	id := "1610612737"
	team, err := datamodel.GetTeam(ctx, id)
	if err != nil {
		fmt.Printf("Error getting team: %v\n", err)
		return
	}

	fmt.Println(*team)
}
