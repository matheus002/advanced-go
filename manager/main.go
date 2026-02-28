package manager

import "errors"

var (
	ErrTruckNotFound = errors.New("Truck not found")
)

type FleeManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tm *truckManager) GetTruck(id string) (*Truck, error) {
	t, ok := tm.trucks[id]

	if !ok {
		return nil, ErrTruckNotFound
	}
	return t, nil
}

func (tm *truckManager) RemoveTruck(id string) error {
	delete(tm.trucks, id)
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	tm.trucks[id].Cargo = cargo
	return nil
}
