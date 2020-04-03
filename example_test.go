package zztermcsi_test

import (
	"fmt"

	"git.sr.ht/~mna/zztermcsi"
)

func ExampleCSI_FuncString() {
	bold := zztermcsi.ChAttr.FuncString(zztermcsi.ChAttrBold)
	norm := zztermcsi.ChAttr.FuncString(zztermcsi.ChAttrNoBoldDim)
	blueBlink := zztermcsi.ChAttr.FuncString(zztermcsi.ChAttrFgBlue, zztermcsi.ChAttrBlink)
	defNoBlink := zztermcsi.ChAttr.FuncString(zztermcsi.ChAttrFgDef, zztermcsi.ChAttrNoBlink)
	fmt.Printf("%sHello%s, %sworld!%s\n", bold, norm, blueBlink, defNoBlink)
}
