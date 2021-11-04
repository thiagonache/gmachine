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
