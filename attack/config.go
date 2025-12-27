package attack

// AttackConfig contains configuration options used to run a brute-force
// simulation. Fields control the mask pattern, the charset used to fill
// mask placeholders, the number of concurrent worker goroutines, and the
// internal channel buffer size for job distribution.
type AttackConfig struct {
	Mask        string
	Charset     string
	WorkerCount int
	BufferSize  int
}
