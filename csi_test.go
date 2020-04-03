package zztermcsi

import (
	"strings"
	"testing"
)

func TestFunc_NoArg(t *testing.T) {
	for i, seq := range csiSeqs {
		t.Run(string(seq), func(t *testing.T) {
			csi := CSI(i)
			want := strings.ReplaceAll(strings.ReplaceAll(string(seq), "\x01", ""), "\x02", "")
			got := csi.FuncString()
			if want != got {
				t.Fatalf("want %q, got %q", want, got)
			}
		})
	}
}

func TestFunc_Args(t *testing.T) {
	cases := []struct {
		csi  CSI
		args []int
		want string
	}{
		{InsCh, nil, "\x1b[@"},
		{InsCh, []int{1}, "\x1b[1@"},
		{InsCh, []int{123}, "\x1b[123@"},
		{InsCh, []int{2, 3, 4}, "\x1b[2@"},
		{CurPos, nil, "\x1b[;H"},
		{CurPos, []int{1}, "\x1b[1;H"},
		{CurPos, []int{1, 2}, "\x1b[1;2H"},
		{CurPos, []int{1, 2, 3}, "\x1b[1;2H"},
		{RstTitleMode, nil, "\x1b[>T"},
		{RstTitleMode, []int{1}, "\x1b[>1T"},
		{RstTitleMode, []int{123, 456}, "\x1b[>123;456T"},
		{RstTitleMode, []int{123, 456, 789}, "\x1b[>123;456;789T"},
		{ChAttrFgIRGB, nil, "\x1b[38;2;;;;m"},
		{ChAttrFgIRGB, []int{1}, "\x1b[38;2;1;;;m"},
		{ChAttrFgIRGB, []int{1, 2, 3}, "\x1b[38;2;1;2;3;m"},
		{ChAttrFgIRGB, []int{1, 2, 3, 4, 5}, "\x1b[38;2;1;2;3;4m"},
		{CSI(255), nil, ""},
	}

	buf := make([]byte, 128)
	for _, c := range cases {
		buf = buf[:0]
		t.Run(c.want, func(t *testing.T) {
			got := string(c.csi.AppendFunc(buf, c.args...))
			if c.want != got {
				t.Fatalf("want %q, got %q", c.want, got)
			}
		})
	}
}

func TestIsCSI(t *testing.T) {
	for _, seq := range csiSeqs {
		if len(seq) == 0 {
			continue
		}

		t.Run(string(seq), func(t *testing.T) {
			if !IsCSI(seq) {
				t.Fatalf("sequence not detected as CSI: %q", seq)
			}
			if !IsCSIString(string(seq)) {
				t.Fatalf("string sequence not detected as CSI: %q", seq)
			}
		})

		// invert the sequence, should NOT be CSI
		iseq := make([]byte, len(seq))
		copy(iseq, seq)
		for i := len(iseq)/2 - 1; i >= 0; i-- {
			opp := len(iseq) - 1 - i
			iseq[i], iseq[opp] = iseq[opp], iseq[i]
		}

		t.Run(string(iseq), func(t *testing.T) {
			if IsCSI(iseq) {
				t.Fatalf("inverted sequence detected as CSI: %q", iseq)
			}
			if IsCSIString(string(iseq)) {
				t.Fatalf("inverted string sequence detected as CSI: %q", iseq)
			}
		})
	}
}
