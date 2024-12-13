package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}
	for _, stage := range stages {
		if stage != nil {
			in = stage(doneStage(in, done))
		}
	}
	return in
}

// stage-wrap for shutdown all operations in a moment.
func doneStage(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out) // without it process will inf waiting
			//nolint:all
			for range in {
				// The goroutine wait that channel has closed. Prevent deadlock and multiple closing
			}
		}()

		// add done-check
		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok { // then in is empty
					return
				}
				out <- val
			}
		}
	}()

	return out
}
