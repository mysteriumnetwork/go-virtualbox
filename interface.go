package virtualbox

import (
	"context"
)

// Manager allows to get and edit every property of Virtualbox.
type Manager interface {
	// Machine gets the machine by its name or UUID
	Machine(context.Context, string) (*Machine, error)

	// ListMachines returns the list of all known machines
	ListMachines(context.Context) ([]*Machine, error)

	// UpdateMachine takes in the properties of the machine and applies the
	// configuration
	UpdateMachine(context.Context, *Machine) error
}
