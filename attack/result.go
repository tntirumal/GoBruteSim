package attack

// AttackResult contains the outcome of an attack run: the discovered
// `Password`, the number of `Attempts` recorded, and the human-readable
// `Duration` the run took.
type AttackResult struct {
	Password string
	Attempts uint64
	Duration string
}
