package csp

// Prefix returns a guarded Process
func Prefix(c Event, m Process) Process {
	return func(e Event) (Process, Event) {
		if e == c {
			return m, e
		}
		return Stop, Bleep
	}
}
