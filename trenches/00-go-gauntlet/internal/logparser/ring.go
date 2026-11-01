package logparser

type RingBuffer struct {
	buf  []Entry
	head int
}

func NewRingBuffer(n int) *RingBuffer {
	return &RingBuffer{buf: make([]Entry, 0, n)}
}

func (r *RingBuffer) Push(e Entry) {
	if len(r.buf) < cap(r.buf) {
		r.buf = append(r.buf, e)
		return
	}
	r.buf[r.head] = e
	r.head = (r.head + 1) % len(r.buf)
}

// Snapshot returns the buffered entries, oldest first.
func (r *RingBuffer) Snapshot() []Entry {
	return r.buf
}

// Last returns the n most recent entries, oldest first.
func (r *RingBuffer) Last(n int) []Entry {
	return r.buf[len(r.buf)-n:]
}
