package main

import (
	"context"
	"log"
)

type contextKey string

var UserIDKey contextKey = "userID"

type Truck struct {
	name  string
	cargo int
}

func main() {
	// truckID := 42
	// anotherTruckID := &truckID
	// log.Println(&truckID)
	// log.Println(*anotherTruckID)

	truck := Truck{cargo: 0}

	fillTruckCargo(&truck)

	log.Printf("%+v\n", truck)
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 43)
	log.Println(ctx)
}

func fillTruckCargo(t *Truck) {
	t.cargo = 100
}
