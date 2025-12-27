package attack

import "sync/atomic"

// AttackStats holds runtime metrics for an attack. The structure uses
// atomic counters to allow concurrent workers to increment attempt counts
// without additional locking.
type AttackStats struct {
	Attempts atomic.Uint64
}
