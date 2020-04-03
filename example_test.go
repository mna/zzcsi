package zzcsi_test

import (
	"fmt"
	"io"

	"git.sr.ht/~mna/zzcsi"
)

func ExampleCSI_FuncString() {
	bold := zzcsi.ChAttr.FuncString(zzcsi.ChAttrBold)
	norm := zzcsi.ChAttr.FuncString(zzcsi.ChAttrNoBoldDim)
	blueBlink := zzcsi.ChAttr.FuncString(zzcsi.ChAttrFgBlue, zzcsi.ChAttrBlink)
	defNoBlink := zzcsi.ChAttr.FuncString(zzcsi.ChAttrFgDef, zzcsi.ChAttrNoBlink)
	fmt.Printf("%sHello%s, %sworld!%s\n", bold, norm, blueBlink, defNoBlink)
}

func ExampleDecodeArgs() {
	// Let's pretend that term is a terminal in raw mode
	var term io.ReadWriter

	// a typical usage would be to send a CSI function to get the cursor
	// position:
	//
	fn := zzcsi.DevStat.Func(zzcsi.DevStatCurPos)
	fmt.Fprint(term, fn)
	//
	// and then get the reply using the zzterm package, e.g.:
	//
	//  for {
	//  	k, err := input.ReadKey(term)
	//  	if err != nil {
	//  		log.Panic(err)
	//  	}
	//
	//  	switch k.Type() {
	//    case zzterm.KeyESCSeq:
	//      b := input.Bytes()
	//  	}
	//  }
	//
	// Let's pretend that's how we got that reply byte sequence:
	b := []byte("\x1b[12;33R")
	var row, col uint64
	n := zzcsi.DecodeArgs(b, &row, &col)
	if n != 2 {
		// oops, did not get both arguments, fail
	}
	// otherwise use row and col
}
