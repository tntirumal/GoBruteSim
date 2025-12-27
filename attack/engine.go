package attack

import (
	"context"
	"miniproject/utils"
	"sync"
	"time"
)

// Run starts the configured attack engine. It launches worker goroutines,
// generates candidate passwords using the configured mask and charset, and
// returns the first successful `AttackResult`. The function uses a
// cancellable context to stop workers once a password is found and collects
// statistics from `AttackStats` for attempts and duration.

type Engine struct {
	Config     AttackConfig
	TargetHash *string
	Stats      *AttackStats
}

func (e *Engine) Run() AttackResult {
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan string, e.Config.BufferSize)
	found := make(chan string)

	wg := sync.WaitGroup{}

	utils.Infof("Starting attack with %d workers", e.Config.WorkerCount)

	for i := 1; i <= e.Config.WorkerCount; i++ {
		wg.Add(1)
		worker := &Worker{
			ID:         i,
			Ctx:        ctx,
			TargetHash: e.TargetHash,
			Jobs:       jobs,
			Found:      found,
			WG:         &wg,
			Stats:      e.Stats,
		}
		go worker.Start()
	}

	go func() {
		for pwd := range utils.GenerateFromMask(e.Config.Mask, e.Config.Charset) {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- pwd
			}
		}
		close(jobs)
	}()

	password := <-found
	cancel()
	wg.Wait()

	return AttackResult{
		Password: password,
		Attempts: e.Stats.Attempts.Load(),
		Duration: time.Since(start).String(),
	}
}
