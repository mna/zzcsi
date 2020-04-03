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

	SetPrvModeAppCurKeys       = 1
	SetPrvModeUSASCIIChSet     = 2
	SetPrvMode132Cols          = 3
	SetPrvModeSmoothScrl       = 4
	SetPrvModeRevVid           = 5
	SetPrvModeOrig             = 6
	SetPrvModeAutoWrap         = 7
	SetPrvModeAutoRep          = 8
	SetPrvModeX10Mous          = 9
	SetPrvModeToolbar          = 10
	SetPrvModeBlinkATT         = 12
	SetPrvModeBlink            = 13
	SetPrvModeXORBlink         = 14
	SetPrvModePrnFF            = 18
	SetPrvModePrnExtFullScr    = 19
	SetPrvModeShowCur          = 25
	SetPrvModeShowScrlbar      = 30
	SetPrvModeFontShFn         = 35
	SetPrvModeTektronix        = 38
	SetPrvModeAllow80_132      = 40
	SetPrvModeMoreFix          = 41
	SetPrvModeNatlReplChSet    = 42
	SetPrvModeMargBl           = 44
	SetPrvModeRevWrap          = 45
	SetPrvModeLog              = 46
	SetPrvModeAltScrBuf        = 47
	SetPrvModeAppKeypad        = 66
	SetPrvModeBackArwBksp      = 67
	SetPrvModeLeftRightMarg    = 69
	SetPrvModeSixelScrl        = 80
	SetPrvModeNoClrScr         = 95
	SetPrvModeX11MousBtn       = 1000
	SetPrvModeX11MousHilite    = 1001
	SetPrvModeX11MousCell      = 1002
	SetPrvModeX11MousAll       = 1003
	SetPrvModeFocus            = 1004
	SetPrvModeUTF8Mous         = 1005
	SetPrvModeSGRMous          = 1006
	SetPrvModeAltScrl          = 1007
	SetPrvModeScrlBotOnOut     = 1010
	SetPrvModeScrlBotOnKey     = 1011
	SetPrvModeURXVTMous        = 1015
	SetPrvModeMetaKey          = 1034
	SetPrvModeAltNumLock       = 1035
	SetPrvModeESCOnMeta        = 1036
	SetPrvModeDELKeypad        = 1037
	SetPrvModeESCOnAlt         = 1039
	SetPrvModeKeepSel          = 1040
	SetPrvModeUseClipbSel      = 1041
	SetPrvModeBlUrg            = 1042
	SetPrvModePopOnBl          = 1043
	SetPrvModeReuseClipb       = 1044
	SetPrvModeSwitchAltScrBuf  = 1046
	SetPrvModeUseAltScrBuf     = 1047
	SetPrvModeSaveCur          = 1048
	SetPrvModeSaveCurAltScrBuf = 1049
	SetPrvModeTinfoTcap        = 1050
	SetPrvModeSun              = 1051
	SetPrvModeHP               = 1052
	SetPrvModeSCO              = 1053
	SetPrvModeLegacyKeyb       = 1060
	SetPrvModeVT220Keyb        = 1061
	SetPrvModeBrackPaste       = 2004

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

	RstPrvModeNormCurKeys       = 1
	RstPrvModeVT52              = 2
	RstPrvMode80Cols            = 3
	RstPrvModeJumpScrl          = 4
	RstPrvModeNormVid           = 5
	RstPrvModeNormCur           = 6
	RstPrvModeNoAutoWrap        = 7
	RstPrvModeNoAutoRep         = 8
	RstPrvModeNoX10Mous         = 9
	RstPrvModeNoToolbar         = 10
	RstPrvModeNoBlinkATT        = 12
	RstPrvModeNoBlink           = 13
	RstPrvModeNoXORBlink        = 14
	RstPrvModeNoPrnFF           = 18
	RstPrvModePrnScrlRegion     = 19
	RstPrvModeHideCur           = 25
	RstPrvModeHideScrlbar       = 30
	RstPrvModeNoFontShFn        = 35
	RstPrvModeDisallow80_132    = 40
	RstPrvModeNoMoreFix         = 41
	RstPrvModeNoNatlReplChSet   = 42
	RstPrvModeNoMargBl          = 44
	RstPrvModeNoRevWrap         = 45
	RstPrvModeNoLog             = 46
	RstPrvModeNormScrBuf        = 47
	RstPrvModeNumKeypad         = 66
	RstPrvModeBackArwDel        = 67
	RstPrvModeNoLeftRightMarg   = 69
	RstPrvModeNoSixelScrl       = 80
	RstPrvModeClrScr            = 95
	RstPrvModeNoX11MousBtn      = 1000
	RstPrvModeNoX11MousHilite   = 1001
	RstPrvModeNoX11MousCell     = 1002
	RstPrvModeNoX11MousAll      = 1003
	RstPrvModeNoFocus           = 1004
	RstPrvModeNoUTF8Mous        = 1005
	RstPrvModeNoSGRMous         = 1006
	RstPrvModeNoAltScrl         = 1007
	RstPrvModeNoScrlBotOnOut    = 1010
	RstPrvModeNoScrlBotOnKey    = 1011
	RstPrvModeNoURXVTMous       = 1015
	RstPrvModeNoMetaKey         = 1034
	RstPrvModeNoAltNumLock      = 1035
	RstPrvModeNoESCOnMeta       = 1036
	RstPrvModeVT220RemKeypad    = 1037
	RstPrvModeNoESCOnAlt        = 1039
	RstPrvModeNoKeepSel         = 1040
	RstPrvModeUsePriSel         = 1041
	RstPrvModeNoBlUrg           = 1042
	RstPrvModeNoPopOnBl         = 1043
	RstPrvModeNoSwitchAltScrBuf = 1046
	RstPrvModeUseNormScrBuf     = 1047
	RstPrvModeRstrCur           = 1048
	RstPrvModeRstrCurNormScrBuf = 1049
	RstPrvModeTinfoTcap         = 1050
	RstPrvModeSun               = 1051
	RstPrvModeHP                = 1052
	RstPrvModeSCO               = 1053
	RstPrvModeLegacyKeyb        = 1060
	RstPrvModeSunPCKeyb         = 1061
	RstPrvModeBrackPaste        = 2004

	ChAttrNorm             = 0
	ChAttrBold             = 1
	ChAttrDim              = 2
	ChAttrItalics          = 3
	ChAttrUnderline        = 4
	ChAttrBlink            = 5
	ChAttrInverse          = 7
	ChAttrHidden           = 8
	ChAttrCrossed          = 9
	ChAttrDoubleUnderline  = 21
	ChAttrNoBoldDim        = 22
	ChAttrNoItalics        = 23
	ChAttrNoUnderline      = 24
	ChAttrNoBlink          = 25
	ChAttrNoInverse        = 27
	ChAttrNoHidden         = 28
	ChAttrNoCrossed        = 29
	ChAttrFgBlack          = 30
	ChAttrFgRed            = 31
	ChAttrFgGreen          = 32
	ChAttrFgYellow         = 33
	ChAttrFgBlue           = 34
	ChAttrFgMagenta        = 35
	ChAttrFgCyan           = 36
	ChAttrFgWhite          = 37
	ChAttrFgDef            = 39
	ChAttrBgBlack          = 40
	ChAttrBgRed            = 41
	ChAttrBgGreen          = 42
	ChAttrBgYellow         = 43
	ChAttrBgBlue           = 44
	ChAttrBgMagenta        = 45
	ChAttrBgCyan           = 46
	ChAttrBgWhite          = 47
	ChAttrBgDef            = 49
	ChAttrFgBlackBright    = 90
	ChAttrFgRedBright      = 91
	ChAttrFgGreenBright    = 92
	ChAttrFgYellowBright   = 93
	ChAttrFgBlueBright     = 94
	ChAttrFgMagentaBright  = 95
	ChAttrFgCyanBright     = 96
	ChAttrFgWhiteBright    = 97
	ChAttrBgBlackBright    = 100
	ChAttrBgRedBright      = 101
	ChAttrBgGreenBright    = 102
	ChAttrBgYellowBright   = 103
	ChAttrBgBlueBright     = 104
	ChAttrBgMagentaBright  = 105
	ChAttrBgCyanBright     = 106
	ChAttrBgWhiteBright    = 107
	ChAttrNo16ColorFgBgDef = 100

	SetKeyModKeyb      = 0
	SetKeyModCurKeys   = 1
	SetKeyModFnKeys    = 2
	SetKeyModOtherKeys = 4

	RstKeyModKeyb      = 0
	RstKeyModCurKeys   = 1
	RstKeyModFnKeys    = 2
	RstKeyModOtherKeys = 4

	DevStatReport = 5
	DevStatCurPos = 6

	DisKeyModKeyb      = 0
	DisKeyModCurKeys   = 1
	DisKeyModFnKeys    = 2
	DisKeyModOtherKeys = 4

	DevStatDECCurPos         = 6
	DevStatDECPrn            = 15
	DevStatDECUDK            = 25
	DevStatDECKeyb           = 26
	DevStatDECLocator        = 53
	DevStatDECLocatorType    = 56
	DevStatDECMacroSpc       = 62
	DevStatDECMemChecksum    = 63
	DevStatDECDataIntegrity  = 75
	DevStatDECMultiSsnConfig = 85

	PtrModeNeverHide             = 0
	PtrModeHideIfNoMousTrack     = 1
	PtrModeAlwaysHideExceptLeave = 2
	PtrModeAlwaysHide            = 3

	SetConfLvlVT100 = 61
	SetConfLvlVT200 = 62
	SetConfLvlVT300 = 63
	SetConfLvlVT400 = 65
	SetConfLvlVT500 = 65
	SetConfLvl8Bit  = 0
	SetConfLvl7Bit  = 1
	SetConfLvl8bit2 = 2

	LdLEDsClrAll     = 0
	LdLEDsNumLock    = 1
	LdLEDsCapsLock   = 2
	LdLEDsScrlLock   = 3
	LdLEDsNoNumLock  = 21
	LdLEDsNoCapsLock = 22
	LdLEDsNoScrlLock = 23

	SetCurStyleBlinkBlock       = 0
	SetCurStyleBlinkBlockDef    = 1
	SetCurStyleNoBlinkBlock     = 2
	SetCurStyleBlinkUnderline   = 3
	SetCurStyleNoBlinkUnderline = 4
	SetCurStyleBlinkBar         = 5
	SetCurStyleNoBlinkBar       = 6

	ChProtAttrCanEraseDef = 0
	ChProtAttrCannotErase = 1
	ChProtAttrCanErase    = 2

	SetTitleModeSetLabelsHex  = 0
	SetTitleModeQryLabelsHex  = 1
	SetTitleModeSetLabelsUTF8 = 2
	SetTitleModeQryLabelsUTF8 = 3

	SetBlVolOff   = 0
	SetBlVolOff2  = 1
	SetBlVolLow   = 2
	SetBlVolLow2  = 3
	SetBlVolLow3  = 4
	SetBlVolHigh  = 5
	SetBlVolHigh2 = 6
	SetBlVolHigh3 = 7
	SetBlVolHigh4 = 8

	SetMargBlVolHigh  = 0
	SetMargBlVolHigh2 = 5
	SetMargBlVolHigh3 = 6
	SetMargBlVolHigh4 = 7
	SetMargBlVolHigh5 = 8
	SetMargBlVolOff   = 1
	SetMargBlVolLow   = 2
	SetMargBlVolLow2  = 3
	SetMargBlVolLow3  = 4
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
	ChAttrFgIRGB
	ChAttrBgIRGB
	ChAttrFgIx
	ChAttrBgIx
	ChAttrFgRGB
	ChAttrBgRGB
	SetKeyMod
	RstKeyMod
	DevStat
	DisKeyMod
	DevStatDEC
	PtrMode
	SoftRst
	SetConfLvl
	ANSIMode
	DECPrvMode
	PushVidAttr
	LdLEDs
	SetCurStyle
	ChProtAttr
	PopVidAttr
	SetScrlRegn
	RstrDECPrvMode
	SetAttrRect
	SaveCur
	LeftRightMarg
	SaveDECPrvMode
	WinOps
	SetTitleMode
	SetBlVol
	RevAttrRect
	RstrCur
	SetMargBlVol
	CopyRect
)

var (
	csiPrefix = []byte("\x1b[")

	// The CSI "Ps" (single number) parameter is encoded as "\x01" and the "Pm"
	// (multiple numbers separated by ;) is encoded as "\x02".

	insCh          = []byte("\x1b[\x01@")
	shLeft         = []byte("\x1b[\x01 @")
	curUp          = []byte("\x1b[\x01A")
	shRight        = []byte("\x1b[\x01 A")
	curDown        = []byte("\x1b[\x01B")
	curFwd         = []byte("\x1b[\x01C")
	curBwd         = []byte("\x1b[\x01D")
	curNextLn      = []byte("\x1b[\x01E")
	curPrevLn      = []byte("\x1b[\x01F")
	curColAbs      = []byte("\x1b[\x01G")
	curPos         = []byte("\x1b[\x01;\x01H")
	curFwdTab      = []byte("\x1b[\x01I")
	eraseScr       = []byte("\x1b[\x01J")
	selEraseScr    = []byte("\x1b[?\x01J")
	eraseLn        = []byte("\x1b[\x01K")
	selEraseLn     = []byte("\x1b[?\x01K")
	insLn          = []byte("\x1b[\x01L")
	delLn          = []byte("\x1b[\x01M")
	delCh          = []byte("\x1b[\x01P")
	scrlUp         = []byte("\x1b[\x01S")
	scrlDown       = []byte("\x1b[\x01T")
	rstTitleMode   = []byte("\x1b[>\x02T")
	eraseCh        = []byte("\x1b[\x01X")
	curBwdTab      = []byte("\x1b[\x01Z")
	chColAbs       = []byte("\x1b[\x02`")
	chColRel       = []byte("\x1b[\x02a")
	repCh          = []byte("\x1b[\x01b")
	priDevAttr     = []byte("\x1b[\x01c")
	terDevAttr     = []byte("\x1b[=\x01c")
	secDevAttr     = []byte("\x1b[>\x01c")
	chLnAbs        = []byte("\x1b[\x02d")
	chLnRel        = []byte("\x1b[\x02e")
	chLnCol        = []byte("\x1b[\x01;\x01f")
	tabClr         = []byte("\x1b[\x01g")
	setMode        = []byte("\x1b[\x02h")
	setPrvMode     = []byte("\x1b[?\x02h")
	mediaCopy      = []byte("\x1b[\x02i")
	mediaCopyDEC   = []byte("\x1b[?\x02i")
	rstMode        = []byte("\x1b[\x02l")
	rstPrvMode     = []byte("\x1b[?\x02l")
	chAttr         = []byte("\x1b[\x02m")
	chAttrFgIRGB   = []byte("\x1b[38;2;\x01;\x01;\x01;\x01m")
	chAttrBgIRGB   = []byte("\x1b[48;2;\x01;\x01;\x01;\x01m")
	chAttrFgIx     = []byte("\x1b[38;5;\x01m")
	chAttrBgIx     = []byte("\x1b[48;5;\x01m")
	chAttrFgRGB    = []byte("\x1b[38;2;\x01;\x01;\x01m")
	chAttrBgRGB    = []byte("\x1b[48;2;\x01;\x01;\x01m")
	setKeyMod      = []byte("\x1b[>\x01;\x01m")
	rstKeyMod      = []byte("\x1b[>\x01m")
	devStat        = []byte("\x1b[\x01n")
	disKeyMod      = []byte("\x1b[>\x02n")
	devStatDEC     = []byte("\x1b[?\x01n")
	ptrMode        = []byte("\x1b[>\x01p")
	softRst        = []byte("\x1b[!p")
	setConfLvl     = []byte("\x1b[\x01;\x01\"p")
	ansiMode       = []byte("\x1b[\x01$p")
	decPrvMode     = []byte("\x1b[?\x01$p")
	pushVidAttr    = []byte("\x1b[\x02#p")
	ldLEDs         = []byte("\x1b[\x01q")
	setCurStyle    = []byte("\x1b[\x01 q")
	chProtAttr     = []byte("\x1b[\x01\"q")
	popVidAttr     = []byte("\x1b[#q")
	setScrlRegn    = []byte("\x1b[\x01;\x01r")
	rstrDECPrvMode = []byte("\x1b[?\x02r")
	setAttrRect    = []byte("\x1b[\x01;\x01;\x01;\x01;\x01$r")
	saveCur        = []byte("\x1b[s")
	leftRightMarg  = []byte("\x1b[\x01;\x01s")
	saveDECPrvMode = []byte("\x1b[?\x02s")
	winOps         = []byte("\x1b[\x01;\x01;\x01t")
	setTitleMode   = []byte("\x1b[>\x02t")
	setBlVol       = []byte("\x1b[\x01 t")
	revAttrRect    = []byte("\x1b[\x01;\x01;\x01;\x01;\x01$t")
	rstrCur        = []byte("\x1b[u")
	setMargBlVol   = []byte("\x1b[\x01 u")
	copyRect       = []byte("\x1b[\x01;\x01;\x01;\x01;\x01;\x01;\x01;\x01$v")
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

	RstTitleMode:   rstTitleMode,
	EraseCh:        eraseCh,
	CurBwdTab:      curBwdTab,
	ChColAbs:       chColAbs,
	ChColRel:       chColRel,
	RepCh:          repCh,
	PriDevAttr:     priDevAttr,
	TerDevAttr:     terDevAttr,
	SecDevAttr:     secDevAttr,
	ChLnAbs:        chLnAbs,
	ChLnRel:        chLnRel,
	ChLnCol:        chLnCol,
	TabClr:         tabClr,
	SetMode:        setMode,
	SetPrvMode:     setPrvMode,
	MediaCopy:      mediaCopy,
	MediaCopyDEC:   mediaCopyDEC,
	RstMode:        rstMode,
	RstPrvMode:     rstPrvMode,
	ChAttr:         chAttr,
	ChAttrFgIRGB:   chAttrFgIRGB,
	ChAttrBgIRGB:   chAttrBgIRGB,
	ChAttrFgIx:     chAttrFgIx,
	ChAttrBgIx:     chAttrBgIx,
	ChAttrFgRGB:    chAttrFgRGB,
	ChAttrBgRGB:    chAttrBgRGB,
	SetKeyMod:      setKeyMod,
	RstKeyMod:      rstKeyMod,
	DevStat:        devStat,
	DisKeyMod:      disKeyMod,
	DevStatDEC:     devStatDEC,
	PtrMode:        ptrMode,
	SoftRst:        softRst,
	SetConfLvl:     setConfLvl,
	ANSIMode:       ansiMode,
	DECPrvMode:     decPrvMode,
	PushVidAttr:    pushVidAttr,
	LdLEDs:         ldLEDs,
	SetCurStyle:    setCurStyle,
	ChProtAttr:     chProtAttr,
	PopVidAttr:     popVidAttr,
	SetScrlRegn:    setScrlRegn,
	RstrDECPrvMode: rstrDECPrvMode,
	SetAttrRect:    setAttrRect,
	SaveCur:        saveCur,
	LeftRightMarg:  leftRightMarg,
	SaveDECPrvMode: saveDECPrvMode,
	WinOps:         winOps,
	SetTitleMode:   setTitleMode,
	SetBlVol:       setBlVol,
	RevAttrRect:    revAttrRect,
	RstrCur:        rstrCur,
	SetMargBlVol:   setMargBlVol,
	CopyRect:       copyRect,
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
