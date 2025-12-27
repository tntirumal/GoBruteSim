package attack

import (
	"context"
	"miniproject/utils"
	"sync"
)

// Start runs the worker loop which receives candidate passwords from the
// `Jobs` channel, computes their hash, updates the shared `Stats`, and sends
// the found password to the `Found` channel when a match occurs. The worker
// respects the provided context to exit early when the attack is cancelled.

type Worker struct {
	ID         int
	Ctx        context.Context
	TargetHash *string
	Jobs       <-chan string
	Found      chan<- string
	WG         *sync.WaitGroup
	Stats      *AttackStats
}

func (w *Worker) Start() {
	defer w.WG.Done()

	utils.Infof("Worker %d started", w.ID)

	for {
		select {
		case <-w.Ctx.Done():
			utils.Warnf("Worker %d stopping", w.ID)
			return

		case pwd, ok := <-w.Jobs:
			if !ok {
				return
			}

			hash := utils.ConvertToHash(&pwd)
			w.Stats.Attempts.Add(1)

			if utils.CompareWithHash(hash, w.TargetHash) {
				utils.Infof("Worker %d FOUND PASSWORD: %s", w.ID, pwd)
				w.Found <- pwd
				return
			}
		}
	}
}
