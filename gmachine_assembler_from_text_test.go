package gmachine_test

import (
	"bytes"
	"gmachine"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestAssembleFromText(t *testing.T) {
// 	t.Parallel()
// 	text := `
// INCA
// CALL 5
// INCA
// #test comment and a blank line

// HALT
// add_one:
// 	INCA
// 	RETN`
// 	g := gmachine.New()
// 	words, err := gmachine.AssembleFromText(text)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	g.RunProgram(words)
// 	var wantA gmachine.Word = 3
// 	if wantA != g.A {
// 		t.Errorf("want initial A value %d, got %d", wantA, g.A)
// 	}
// 	var wantP gmachine.Word = 5
// 	if wantP != g.P {
// 		t.Errorf("want initial P value %d, got %d", wantP, g.P)
// 	}
// 	var wantN gmachine.Word = 0
// 	if wantN != g.N {
// 		t.Errorf("want initial N value %d, got %d", wantN, g.N)
// 	}
// }

func TestHelloWorld(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	words, err := gmachine.AssembleFromText(`
		JUMP 12
		72 101 108 108 111 87 111 114 108 100
		SETI 2
		SETA [I]
		BIOS IOWRITE STDOUT
		INCI
		CMPI 12
		JEQ 14
	`)
	if err != nil {
		t.Fatal(err)
	}
	buf := &bytes.Buffer{}
	g.Stdout = buf
	g.RunProgram(words)
	want := "HelloWorld"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestHelloWorldRune(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	words, err := gmachine.AssembleFromText(`
		JUMP 12
		"HelloWorld"
		SETI 2
		SETA [I]
		BIOS IOWRITE STDOUT
		INCI
		CMPI 12
		JEQ 14
	`)
	if err != nil {
		t.Fatal(err)
	}
	buf := &bytes.Buffer{}
	g.Stdout = buf
	g.RunProgram(words)
	want := "HelloWorld"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestTestAssembleData(t *testing.T) {
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
func TestAssembleData(t *testing.T) {
	t.Parallel()
	want := []gmachine.Word{65, 98, 99}
	got, err := gmachine.AssembleData(`"Abc"`)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
