package counting

import "io"

type Writer struct {
	w io.Writer
	c int
}

func NewWriter(w io.Writer) (h *Writer) {
	h = new(Writer)
	h.w = w
	h.c = 0
	return
}

func (h *Writer) Write(p []byte) (n int, err error) {
	n, err = h.w.Write(p)
	h.c += n
	return
}

// Count: returns the total number of bytes that were 
// written to the underlying writer since the last call 
// to Clear
func (h *Writer) Count() int {
	return h.c
}

// Clear: zeroes the write byte counter
func (h *Writer) Clear() {
	h.c = 0
}
