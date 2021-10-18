package gmachine_test

import (
	"bytes"
	"gmachine"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRunProgram(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{
		gmachine.NOOP,
		gmachine.HALT,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestSubstract2From3(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Memory[1] = gmachine.INCA
	g.Memory[2] = gmachine.INCA
	g.Memory[3] = gmachine.DECA
	g.Memory[4] = gmachine.DECA
	g.Run()
	var wantA gmachine.Word = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestSubstractTwo(t *testing.T) {
	testCases := []struct {
		desc                 string
		valueA, wantA, wantP gmachine.Word
	}{
		{
			desc:   "Substract 2 from 3",
			valueA: 3,
			wantA:  1,
			wantP:  5,
		},
		{
			desc:   "Substract 2 from 200",
			valueA: 200,
			wantA:  198,
			wantP:  5,
		},
	}
	for _, tC := range testCases {
		g := gmachine.New()
		t.Run(tC.desc, func(t *testing.T) {
			g.Memory[0] = gmachine.SETA
			g.Memory[1] = tC.valueA
			g.Memory[2] = gmachine.DECA
			g.Memory[3] = gmachine.DECA
			g.Run()
			if tC.wantA != g.A {
				t.Errorf("want A value %d, got %d", tC.wantA, g.A)
			}
			if tC.wantP != g.P {
				t.Errorf("want P value %d, got %d", tC.wantP, g.P)
			}
		})
	}
}

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

func TestReadWords(t *testing.T) {
	t.Parallel()
	want := []gmachine.Word{gmachine.SETA, math.MaxUint64, gmachine.DECA}
	input := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		255, 255, 255, 255, 255, 255, 255, 255,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	})
	got, err := gmachine.ReadWords(input)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWriteWords(t *testing.T) {
	t.Parallel()
	input := []gmachine.Word{gmachine.SETA, 10, gmachine.DECA}
	want := []byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		0, 0, 0, 0, 0, 0, 0, 10,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	}
	output := &bytes.Buffer{}
	err := gmachine.WriteWords(output, input)
	if err != nil {
		t.Fatal(err)
	}
	got := output.Bytes()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
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
	var wantL gmachine.Word = 2
	if wantL != g.L {
		t.Errorf("want initial L value %d, got %d", wantL, g.L)
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
	var wantL gmachine.Word = 0
	if wantL != g.L {
		t.Errorf("want initial L value %d, got %d", wantL, g.L)
	}
}

func TestExecuteBinary(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	err := g.ExecuteBinary("testdata/setadeca.gbin")
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 3
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}
