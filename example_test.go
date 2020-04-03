package zzcsi_test

import (
	"fmt"

	"git.sr.ht/~mna/zzcsi"
)

func ExampleCSI_FuncString() {
	bold := zzcsi.ChAttr.FuncString(zzcsi.ChAttrBold)
	norm := zzcsi.ChAttr.FuncString(zzcsi.ChAttrNoBoldDim)
	blueBlink := zzcsi.ChAttr.FuncString(zzcsi.ChAttrFgBlue, zzcsi.ChAttrBlink)
	defNoBlink := zzcsi.ChAttr.FuncString(zzcsi.ChAttrFgDef, zzcsi.ChAttrNoBlink)
	fmt.Printf("%sHello%s, %sworld!%s\n", bold, norm, blueBlink, defNoBlink)
}
