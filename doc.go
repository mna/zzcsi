// Package zzcsi implements functions for the execution of Control Sequence
// Introducer (CSI) sequences on xterm-compatible terminals. The complete
// documentation of CSI functions is available at [1].
//
// Usage
//
// All CSI functions are represented by the CSI type and a corresponding constant.
// Call Func to get the byte sequence to write to the terminal to execute a
// given CSI function. FuncString is the same but returns a string that can be
// easily inserted e.g. in a fmt.Fprintf call (for example, to wrap some text
// in a bold or colored attribute before resetting it to normal).
//
// For efficient generation of CSI sequences, use AppendFunc to append the bytes
// to a caller-provided buffer.
//
//     [1]: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Functions-using-CSI-_-ordered-by-the-final-character_s_
//
package zzcsi // import "git.sr.ht/~mna/zzcsi"
