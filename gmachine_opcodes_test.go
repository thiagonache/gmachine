package gmachine_test

import (
	"bytes"
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP gmachine.Word = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue gmachine.Word = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
	var wantA gmachine.Word = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantE gmachine.Word = 0
	if wantA != g.E {
		t.Errorf("want initial A value %d, got %d", wantE, g.E)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Run()
	var wantP gmachine.Word = 1
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.NOOP
	g.Run()
	var wantP gmachine.Word = 2
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Run()
	var wantA gmachine.Word = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.A = 2
	g.Memory[0] = gmachine.DECA
	g.Run()
	var wantA gmachine.Word = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestBIOS(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.SETA
	g.Memory[1] = 'A'
	g.Memory[2] = gmachine.BIOS
	g.Memory[3] = gmachine.IOWrite
	g.Memory[4] = gmachine.PortStdout
	buf := &bytes.Buffer{}
	g.Stdout = buf
	g.Run()
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

func TestBIOSStderr(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.SETA
	g.Memory[1] = 'A'
	g.Memory[2] = gmachine.BIOS
	g.Memory[3] = gmachine.IOWrite
	g.Memory[4] = gmachine.PortStderr
	buf := &bytes.Buffer{}
	g.Stderr = buf
	g.Run()
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

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.SETA
	g.Memory[1] = 5
	g.Run()
	var wantA gmachine.Word = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestJUMP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Memory[1] = gmachine.JUMP
	g.Memory[2] = 0
	g.Memory[3] = gmachine.JUMP
	g.Memory[4] = 0
	g.Memory[5] = gmachine.DECA
	g.Run()

	var wantA gmachine.Word = 2
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}
