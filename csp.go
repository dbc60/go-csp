package csp

// An Event is an element of the alphabet for a process
type Event uint

// An alphabet for some simple processes.
const (
	// Bleep represents an event for a broken or stopped process
	Bleep = Event(iota)
)

// String returns a string representation of an Event.
func (e Event) String() string {
	switch e {
	case Bleep:
		return "Bleep"
	default:
		return "unknown event"
	}
}

// A Process is a function that returns a Process and an Event.
type Process func(e Event) (Process, Event)

// Stop returns itself and the Bleep event, regardless of the value of e.
func Stop(e Event) (Process, Event) {
	return Stop, Bleep
}
