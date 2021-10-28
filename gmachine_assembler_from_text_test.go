package gmachine_test

import (
	"bytes"
	"gmachine"
	"testing"
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
		SETA 65
		BIOS IOWRITE STDOUT
	`)
	if err != nil {
		t.Fatal(err)
	}
	buf := &bytes.Buffer{}
	g.Stdout = buf
	g.RunProgram(words)
	want := "A"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
