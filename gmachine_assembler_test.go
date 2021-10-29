package gmachine_test

import (
	"gmachine"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAssemble(t *testing.T) {
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
	input := []string{""}
	_, err := gmachine.Assemble(input)
	if err == nil {
		t.Errorf("An error is expected but not found")
	}
}

func TestAssembleOperand(t *testing.T) {
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

func TestAssembleOperandInvalid(t *testing.T) {
	t.Parallel()
	input := []string{"SETA", "DECA"}
	_, err := gmachine.Assemble(input)
	if err == nil {
		t.Error("Expecting error but not found")
	}
}
