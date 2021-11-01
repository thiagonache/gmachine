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
	wantZ := false
	if wantZ != g.FlagZ {
		t.Errorf("want initial A value %t, got %t", wantZ, g.FlagZ)
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

func TestINCI(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCI
	g.Run()
	var wantI gmachine.Word = 1
	if wantI != g.I {
		t.Errorf("want initial A value %d, got %d", wantI, g.I)
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

func TestCMPA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.CMPA
	g.Memory[1] = 5
	g.Run()
	wantZ := false
	if wantZ != g.FlagZ {
		t.Errorf("want flag Z value %t, got %t", wantZ, g.FlagZ)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestSETI(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.SETI
	g.Memory[1] = 5
	g.Run()
	var wantI gmachine.Word = 5
	if wantI != g.I {
		t.Errorf("want initial A value %d, got %d", wantI, g.I)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestCMPI(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.CMPI
	g.Memory[1] = 5
	g.Run()
	wantZ := false
	if wantZ != g.FlagZ {
		t.Errorf("want flag Z value %t, got %t", wantZ, g.FlagZ)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestCMPASetZ(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Memory[1] = gmachine.INCA
	g.Memory[2] = gmachine.CMPA
	g.Memory[3] = 2
	g.Run()
	wantZ := true
	if wantZ != g.FlagZ {
		t.Errorf("want flag Z value %t, got %t", wantZ, g.FlagZ)
	}
	var wantP gmachine.Word = 5
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}
