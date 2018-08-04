package csp

// Choice2 returns a Process that will "choose" to return one of P, Q, or Stop
// depending on Event e matching c, d, or neither.
func Choice2(c Event, p Process, d Event, q Process) Process {
	return func(e Event) (Process, Event) {
		if e == c {
			return p, e
		}
		if e == d {
			return q, e
		}
		return Stop, Bleep
	}
}
