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

func TestHelloWorldString(t *testing.T) {
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

func TestBREAK(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	words, err := gmachine.AssembleFromText(`
		INCA
		BREAK
		INCA
		HALT
	`)
	if err != nil {
		t.Fatal(err)
	}
	g.RunProgram(words)
	// at breakpoint
	want := gmachine.State{
		P:    2,
		A:    1,
		I:    0,
		Z:    false,
		Next: gmachine.INCA,
	}
	got := g.State()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	// run to end
	g.Run()
	want = gmachine.State{
		P:    4,
		A:    2,
		I:    0,
		Z:    false,
		Next: gmachine.HALT,
	}
	got = g.State()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStateString(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	words, err := gmachine.AssembleFromText(`
		INCA
		HALT
	`)
	if err != nil {
		t.Fatal(err)
	}
	g.RunProgram(words)
	want := "P: 2 A: 1 I: 0 Z: false Next: 0"
	got := g.State().String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
