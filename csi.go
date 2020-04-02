package zztermcsi

import (
	"bytes"
	"strconv"
	"strings"
)

// CSI represents a Control Sequence Introducer function as supported
// by xterm-compatible terminals.
//
// See https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Functions-using-CSI-_-ordered-by-the-final-character_s_
// for details.
type CSI byte

// List of CSI functions.
const (
	InsCh CSI = iota
	ShLeft
	CurUp
	ShRight
	CurDown
	CurFwd
	CurBwd
)

var (
	csiPrefix = []byte("\x1b[")

	// The CSI "Ps" (single number) parameter is encoded as "\x01" and the "Pm"
	// (multiple numbers separated by ;) is encoded as "\x02".

	insCh   = []byte("\x1b[\x01@")
	shLeft  = []byte("\x1b[\x01 @")
	curUp   = []byte("\x1b[\x01A")
	shRight = []byte("\x1b[\x01 A")
	curDown = []byte("\x1b[\x01B")
	curFwd  = []byte("\x1b[\x01C")
	curBwd  = []byte("\x1b[\x01D")
)

var csiSeqs = [...][]byte{
	InsCh:   insCh,
	ShLeft:  shLeft,
	CurUp:   curUp,
	ShRight: shRight,
	CurDown: curDown,
	CurFwd:  curFwd,
	CurBwd:  curBwd,
}

// Func returns the sequence of bytes to execute this CSI function with
// the provided numeric arguments. Note that no validation is done regarding
// the number of arguments - if the function supports a single argument, only
// one will be inserted, if it supports many, all expected arguments will be
// inserted. If less arguments than those expected are provided, the remaining
// arguments are left unspecified (which usually results in a default value
// fallback).
func (c CSI) Func(args ...int) []byte {
	if int(c) >= len(csiSeqs) {
		return nil
	}
	seq := csiSeqs[c]
	buf := make([]byte, 0, len(seq))
	return appendFunc(buf, seq, args)
}

// FuncString is like Func except it returns a string value. This can be useful
// to insert e.g. in a printf-style string.
func (c CSI) FuncString(args ...int) string {
	return string(c.Func(args...))
}

// AppendFunc is like Func except it appends the resulting sequence of bytes to
// b and returns the new slice. If b has a large enough capacity to hold the
// sequence, no allocation is made.
func (c CSI) AppendFunc(b []byte, args ...int) []byte {
	if int(c) >= len(csiSeqs) {
		return nil
	}
	seq := csiSeqs[c]
	return appendFunc(b, seq, args)
}

func appendFunc(buf, seq []byte, args []int) []byte {
	// start by processing the Pm (multiple numbers separated by ;), as there
	// can be only one placeholder if Pm is used.
	if ix := bytes.IndexByte(seq, '\x02'); ix >= 0 {
		buf = append(buf, seq[:ix]...)
		for i, arg := range args {
			if i > 0 {
				buf = append(buf, ';')
			}
			buf = strconv.AppendInt(buf, int64(arg), 10)
		}
		buf = append(buf, seq[ix+1:]...)
		return buf
	}

	// otherwise replace the Ps (single number) placeholders, there can be many.
	start := 0
	for start < len(seq) {
		ix := bytes.IndexByte(seq[start:], '\x01')
		if ix < 0 {
			buf = append(buf, seq[start:]...)
			break
		}
		buf = append(buf, seq[start:ix]...)
		start = ix + 1
		if len(args) > 0 {
			buf = strconv.AppendInt(buf, int64(args[0]), 10)
			args = args[1:]
		}
	}
	return buf
}

// IsCSI returns true if b starts with the Control Sequence Introducer
// bytes ("\x1b[", or <ESC> followed by '[').
func IsCSI(b []byte) bool {
	return bytes.HasPrefix(b, csiPrefix)
}

// IsCSIString returns true if s starts with the Control Sequence Introducer
// prefix (see IsCSI for details).
func IsCSIString(s string) bool {
	return strings.HasPrefix(s, string(csiPrefix))
}