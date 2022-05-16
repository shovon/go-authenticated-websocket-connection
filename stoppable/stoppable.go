// Package stoppable holds a single struct that can be used for stopping a
// process.
package stoppable

import "sync"

// Stoppable is an object that we can inherit to represent a stopped process.
//
// Great for processes that are long running but needs to be stopped on command.
//
// Must be created using the `NewStoppable` function!
type Stoppable struct {
	once    *sync.Once
	stopped chan struct{}
}

// NewStoppable creates a new Stoppable instance
func NewStoppable() Stoppable {
	return Stoppable{&sync.Once{}, make(chan struct{})}
}

// OnStopped returns a channel where it will close once the stoppable has been
// stopped.
//
// Usage:
//
//     stopped := s.OnStopped()
//     <-stopped
func (s Stoppable) OnStopped() <-chan struct{} {
	return s.stopped
}

// Stop stops the stoppable.
func (s Stoppable) Stop() {
	s.once.Do(func() {
		close(s.stopped)
	})
}
