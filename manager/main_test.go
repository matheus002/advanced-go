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

func TestConcurrentUpdate(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	const numGoroutines = 100
	const iterations = 100

	done := make(chan bool)

	for i := 0; i < iterations; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				truck, _ := manager.GetTruck("1")
				manager.UpdateTruckCargo("1", truck.Cargo+1)
			}
			done <- true
		}()
	}

	for range numGoroutines {
		<-done
	}

	expectedFinalValue := numGoroutines*iterations + 100
	finalTruck, _ := manager.GetTruck("1")

	if finalTruck.Cargo != expectedFinalValue {
		t.Errorf("Expected final cargo to be %d, but got %d. Race condition detected!", expectedFinalValue, finalTruck.Cargo)
	}
}
