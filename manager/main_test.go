package manager

import "testing"

func TestAddTruck(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	if len(manager.trucks) != 1 {
		t.Errorf("Expected 1 truck, got %d", len(manager.trucks))
	}
}

func TestGetTruck(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	truck, err := manager.GetTruck("1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if truck.ID != "1" {
		t.Errorf("Expected truck ID to be 1, got %s", truck.ID)
	}
}
