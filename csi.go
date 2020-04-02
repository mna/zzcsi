package zztermcsi

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	EraseScrBelow    = 0
	EraseScrAbove    = 1
	EraseScrAll      = 2
	EraseScrSavedLns = 3

	SelEraseScrBelow    = 0
	SelEraseScrAbove    = 1
	SelEraseScrAll      = 2
	SelEraseScrSavedLns = 3

	EraseLnRight = 0
	EraseLnLeft  = 1
	EraseLnAll   = 2

	SelEraseLnRight = 0
	SelEraseLnLeft  = 1
	SelEraseLnAll   = 2

	RstTitleModeSetLabelsHex  = 0
	RstTitleModeQryLabelsHex  = 1
	RstTitleModeSetLabelsUTF8 = 2
	RstTitleModeQryLabelsUTF8 = 3

	TabClrCurrentCol = 0
	TabClrAll        = 3

	SetModeKeybAction = 2
	SetModeIns        = 4
	SetModeSendRecv   = 12
	SetModeAutoNewln  = 20

	// TODO: SetPrvMode constants

	MediaCopyPrnScr      = 0
	MediaCopyPrnContrOff = 4
	MediaCopyPrnContrOn  = 5
	MediaCopyHTMLScr     = 10
	MediaCopySVGScr      = 11

	MediaCopyDECPrnCurLn    = 1
	MediaCopyDECAutoPrnOff  = 4
	MediaCopyDECAutoPrnOn   = 5
	MediaCopyDECPrnCompDisp = 10
	MediaCopyDECPrnAll      = 11

	RstModeKeybAction = 2
	RstModeRepl       = 4
	RstModeSendRecv   = 12
	RstModeNormNewln  = 20

	// TODO: RstPrvMode constants

	// TODO: ChAttr constants

	SetKeyModKeyb      = 0
	SetKeyModCurKeys   = 1
	SetKeyModFnKeys    = 2
	SetKeyModOtherKeys = 4

	RstKeyModKeyb      = 0
	RstKeyModCurKeys   = 1
	RstKeyModFnKeys    = 2
	RstKeyModOtherKeys = 4
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
	CurNextLn
	CurPrevLn
	CurColAbs
	CurPos
	CurFwdTab
	EraseScr
	SelEraseScr
	EraseLn
	SelEraseLn
	InsLn
	DelLn
	DelCh
	ScrlUp
	_ // TODO: Set or request graphics attribute
	ScrlDown
	_ // TODO: Initiate highlight mouse tracking
	RstTitleMode
	EraseCh
	CurBwdTab
	ChColAbs
	ChColRel
	RepCh
	PriDevAttr
	TerDevAttr
	SecDevAttr
	ChLnAbs
	ChLnRel
	ChLnCol
	TabClr
	SetMode
	SetPrvMode
	MediaCopy
	MediaCopyDEC
	RstMode
	RstPrvMode
	ChAttr
	SetKeyMod
	RstKeyMod
)

var (
	csiPrefix = []byte("\x1b[")

	// The CSI "Ps" (single number) parameter is encoded as "\x01" and the "Pm"
	// (multiple numbers separated by ;) is encoded as "\x02".

	insCh        = []byte("\x1b[\x01@")
	shLeft       = []byte("\x1b[\x01 @")
	curUp        = []byte("\x1b[\x01A")
	shRight      = []byte("\x1b[\x01 A")
	curDown      = []byte("\x1b[\x01B")
	curFwd       = []byte("\x1b[\x01C")
	curBwd       = []byte("\x1b[\x01D")
	curNextLn    = []byte("\x1b[\x01E")
	curPrevLn    = []byte("\x1b[\x01F")
	curColAbs    = []byte("\x1b[\x01G")
	curPos       = []byte("\x1b[\x01;\x01H")
	curFwdTab    = []byte("\x1b[\x01I")
	eraseScr     = []byte("\x1b[\x01J")
	selEraseScr  = []byte("\x1b[?\x01J")
	eraseLn      = []byte("\x1b[\x01K")
	selEraseLn   = []byte("\x1b[?\x01K")
	insLn        = []byte("\x1b[\x01L")
	delLn        = []byte("\x1b[\x01M")
	delCh        = []byte("\x1b[\x01P")
	scrlUp       = []byte("\x1b[\x01S")
	scrlDown     = []byte("\x1b[\x01T")
	rstTitleMode = []byte("\x1b[>\x02T")
	eraseCh      = []byte("\x1b[\x01X")
	curBwdTab    = []byte("\x1b[\x01Z")
	chColAbs     = []byte("\x1b[\x02`")
	chColRel     = []byte("\x1b[\x02a")
	repCh        = []byte("\x1b[\x01b")
	priDevAttr   = []byte("\x1b[\x01c")
	terDevAttr   = []byte("\x1b[=\x01c")
	secDevAttr   = []byte("\x1b[>\x01c")
	chLnAbs      = []byte("\x1b[\x02d")
	chLnRel      = []byte("\x1b[\x02e")
	chLnCol      = []byte("\x1b[\x01;\x01f")
	tabClr       = []byte("\x1b[\x01g")
	setMode      = []byte("\x1b[\x02h")
	setPrvMode   = []byte("\x1b[?\x02h")
	mediaCopy    = []byte("\x1b[\x02i")
	mediaCopyDEC = []byte("\x1b[?\x02i")
	rstMode      = []byte("\x1b[\x02l")
	rstPrvMode   = []byte("\x1b[?\x02l")
	chAttr       = []byte("\x1b[\x02m")
	setKeyMod    = []byte("\x1b[>\x01;\x01m")
	rstKeyMod    = []byte("\x1b[>\x01m")
)

var csiSeqs = [...][]byte{
	InsCh:       insCh,
	ShLeft:      shLeft,
	CurUp:       curUp,
	ShRight:     shRight,
	CurDown:     curDown,
	CurFwd:      curFwd,
	CurBwd:      curBwd,
	CurNextLn:   curNextLn,
	CurPrevLn:   curPrevLn,
	CurColAbs:   curColAbs,
	CurPos:      curPos,
	CurFwdTab:   curFwdTab,
	EraseScr:    eraseScr,
	SelEraseScr: selEraseScr,
	EraseLn:     eraseLn,
	SelEraseLn:  selEraseLn,
	InsLn:       insLn,
	DelLn:       delLn,
	DelCh:       delCh,
	ScrlUp:      scrlUp,

	ScrlDown: scrlDown,

	RstTitleMode: rstTitleMode,
	EraseCh:      eraseCh,
	CurBwdTab:    curBwdTab,
	ChColAbs:     chColAbs,
	ChColRel:     chColRel,
	RepCh:        repCh,
	PriDevAttr:   priDevAttr,
	TerDevAttr:   terDevAttr,
	SecDevAttr:   secDevAttr,
	ChLnAbs:      chLnAbs,
	ChLnRel:      chLnRel,
	ChLnCol:      chLnCol,
	TabClr:       tabClr,
	SetMode:      setMode,
	SetPrvMode:   setPrvMode,
	MediaCopy:    mediaCopy,
	MediaCopyDEC: mediaCopyDEC,
	RstMode:      rstMode,
	RstPrvMode:   rstPrvMode,
	ChAttr:       chAttr,
	SetKeyMod:    setKeyMod,
	RstKeyMod:    rstKeyMod,
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
		ix += start
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
