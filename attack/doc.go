// Package attack implements a simple concurrent brute-force password attack
// simulation used by the `miniproject` example. It provides the core types
// and components to configure and run an attack:
//
// - AttackConfig: configuration for mask, charset and worker parameters.
// - Engine: orchestrates worker goroutines and coordinates job distribution.
// - Worker: performs password hashing and comparison against the target.
// - AttackStats / AttackResult: hold runtime metrics and the final result.
//
// The package is intentionally small and demonstrates patterns for
// concurrent workers, cancellation via context, and channel-based
// communication. It is designed for educational/demo usage rather than
// production-grade security tooling.
package attack
