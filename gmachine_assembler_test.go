package gmachine_test

import (
	"gmachine"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAssembleFromSlice(t *testing.T) {
	t.Parallel()

	input := []string{"HALT", "NOOP"}
	want := []gmachine.Word{gmachine.HALT, gmachine.NOOP}
	got, err := gmachine.Assemble(input)
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestAssembleInvalid(t *testing.T) {
	t.Parallel()
	_, err := gmachine.AssembleFromText("")
	if err == nil {
		t.Errorf("An error is expected but not found")
	}
}

func TestAssembleFromSliceOperand(t *testing.T) {
	t.Parallel()
	input := []string{"SETA", "5"}
	want := []gmachine.Word{gmachine.SETA, 5}
	got, err := gmachine.Assemble(input)
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssembleFromSliceMissingOperand(t *testing.T) {
	t.Parallel()
	input := []string{"SETA", "DECA"}
	_, err := gmachine.Assemble(input)
	if err == nil {
		t.Error("Expecting error but not found")
	}
}

func TestAssembleData(t *testing.T) {
	testCases := []struct {
		code, desc string
		want       []gmachine.Word
	}{
		{
			code: `'A'`,
			desc: "Assemble string 'A'",
			want: []gmachine.Word{65},
		},
		{
			code: `"A"`,
			desc: "Assemble string \"A\"",
			want: []gmachine.Word{65},
		},
		{
			code: `"Abc"`,
			desc: "Assemble string \"Abc\"",
			want: []gmachine.Word{65, 98, 99},
		},
		{
			code: `90`,
			desc: "Assemble ASCII decimal 90 (Z)",
			want: []gmachine.Word{90},
		},
		{
			code: `120 121 122`,
			desc: "Assemble ASCII decimals 120 121 122 (xyz)",
			want: []gmachine.Word{120, 121, 122},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := gmachine.AssembleData(tC.code)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(tC.want, got) {
				t.Error(cmp.Diff(tC.want, got))
			}
		})
	}
}

func TestAssembleOperand(t *testing.T) {
	constants := gmachine.PredefinedConstants
	testCases := []struct {
		code, desc string
		want       gmachine.Word
	}{
		{
			code: "2",
			desc: "Assemble decimal 2",
			want: gmachine.Word(2),
		},
		{
			code: "10",
			desc: "Assemble decimal 10",
			want: gmachine.Word(10),
		},
		{
			code: "[I]",
			desc: "Assemble dereference [I]",
			want: gmachine.SETAM,
		},
		{
			code: "IOWRITE",
			desc: "Assemble constant IOWrite",
			want: gmachine.IOWrite,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := gmachine.AssembleOperand(constants, tC.code)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(tC.want, got) {
				t.Error(cmp.Diff(tC.want, got))
			}
		})
	}
}
