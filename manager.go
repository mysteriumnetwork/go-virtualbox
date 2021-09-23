package virtualbox

import (
	"context"
	"sync"
)

var defaultManager = NewManager()

// manager implements all the functionality of the Manager, and is the default
// one used.
type manager struct {
	// Wrap around the existing code until its migrated
	cmd Command
	// lock the whole manager to only allow one action at a time
	// TODO: Decide is this a good idea, or should we have one mutex per
	//       type of operation
	lock sync.Mutex
}

// NewManager returns the real instance of the manager
func NewManager() *manager {
	return &manager{
		cmd: Manage(),
	}
}

// run is the internal function used by other commands.
func (m *manager) run(ctx context.Context, args ...string) (string, string, error) {
	return m.cmd.runOutErrContext(ctx, args...)
}

// Run is a helper function using the defaultManager and can be used to directly
// run commands which are not exposed as part of the Manager API. It returns the
// stdout, stderr and any errors which happened while executing the command.
// The `VBoxManage` argument should not be specified at the beginning as it is
// deducted from the environment.
//
// Notice: Its possible that if we ever cover the API 1:1, this function might
//         be deprecated and later removed.
func Run(ctx context.Context, args ...string) (string, string, error) {
	return defaultManager.run(ctx, args...)
}
