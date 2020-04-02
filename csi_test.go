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
