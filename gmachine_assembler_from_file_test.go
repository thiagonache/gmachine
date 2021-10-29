package gmachine_test

import (
	"bytes"
	"gmachine"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAssembleFromFile(t *testing.T) {
	t.Parallel()

	want := []gmachine.Word{gmachine.HALT}
	got, err := gmachine.AssembleFromFile("testdata/local.gasm")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssembleFromFileSetA(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/seta.gasm")
	if err != nil {
		t.Fatal(err)
	}
	g := gmachine.New()
	g.RunProgram(words)
	var wantA gmachine.Word = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestAssembleFromFileSetADeca(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/setadeca.gasm")
	if err != nil {
		t.Fatal(err)
	}
	g := gmachine.New()
	g.RunProgram(words)
	var wantA gmachine.Word = 3
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 5
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestAssembleFromFileLowercase(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/lowercase.gasm")
	if err != nil {
		t.Fatal(err)
	}
	g := gmachine.New()
	g.RunProgram(words)
	var wantA gmachine.Word = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestAssembleFromFileMultipleOperands(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/biosstdout.gasm")
	if err != nil {
		t.Fatal(err)
	}
	g := gmachine.New()
	buf := &bytes.Buffer{}
	g.Stdout = buf
	g.RunProgram(words)
	want := "A"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	var wantP gmachine.Word = 6
	gotP := g.P
	if wantP != gotP {
		t.Errorf("wantP %d, got %d", wantP, gotP)
	}
}

func TestAssembleToFile(t *testing.T) {
	t.Parallel()
	outPath := filepath.Join(t.TempDir(), "setadeca.gbin")
	wantPath := "testdata/setadeca.gbin"
	err := gmachine.AssembleFromFileToBinary("testdata/setadeca.gasm", outPath)
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
