package gmachine_test

import (
	"bytes"
	"gmachine"
	"testing"
)

func TestRunProgramFromReader(t *testing.T) {
	t.Parallel()
	// SETA 258
	// DECA
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		0, 0, 0, 0, 0, 0, 1, 2,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 257
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}

	var wantP gmachine.Word = 4
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestJEQFromReader(t *testing.T) {
	t.Parallel()
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.INCA,
		0, 0, 0, 0, 0, 0, 0, gmachine.CMPA,
		0, 0, 0, 0, 0, 0, 0, 10,
		0, 0, 0, 0, 0, 0, 0, gmachine.JEQ,
		0, 0, 0, 0, 0, 0, 0, 0,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 10
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}

	var wantP gmachine.Word = 6
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestJUMPFromReader(t *testing.T) {
	t.Parallel()
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.JUMP,
		0, 0, 0, 0, 0, 0, 0, 10,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}

	var wantP gmachine.Word = 11
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestCALLFromReader(t *testing.T) {
	t.Parallel()
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.INCA,
		0, 0, 0, 0, 0, 0, 0, gmachine.CALL,
		0, 0, 0, 0, 0, 0, 0, 10,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 11
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantN gmachine.Word = 3
	if wantN != g.N {
		t.Errorf("want initial N value %d, got %d", wantN, g.N)
	}
}

func TestRETNFromReader(t *testing.T) {
	t.Parallel()
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.INCA,
		0, 0, 0, 0, 0, 0, 0, gmachine.CALL,
		0, 0, 0, 0, 0, 0, 0, 5,
		0, 0, 0, 0, 0, 0, 0, gmachine.INCA,
		0, 0, 0, 0, 0, 0, 0, gmachine.HALT,
		0, 0, 0, 0, 0, 0, 0, gmachine.INCA,
		0, 0, 0, 0, 0, 0, 0, gmachine.RETN,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 3
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 5
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantN gmachine.Word = 0
	if wantN != g.N {
		t.Errorf("want initial N value %d, got %d", wantN, g.N)
	}
}
