// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	HALT = iota
	NOOP
	INCA
	DECA
	SETA
	BIOS
	CMPA
	JEQ
	JUMP
	CALL
	RETN
	INCI
	CMPI
	SETI
	SETAM
)

const (
	IOWrite = iota
	IORead
)

const (
	PortStdin = iota
	PortStdout
	PortStderr
)

var PredefinedConstants = map[string]Word{
	"IOWRITE": IOWrite,
	"IOREAD":  IORead,
	"STDIN":   PortStdin,
	"STDOUT":  PortStdout,
	"STDERR":  PortStderr,
}

type Instruction struct {
	Opcode   Word
	Operands int
}

var TranslateTable = map[string]Instruction{
	"HALT": {Opcode: HALT, Operands: 0},
	"NOOP": {Opcode: NOOP, Operands: 0},
	"SETA": {Opcode: SETA, Operands: 1},
	"DECA": {Opcode: DECA, Operands: 0},
	"INCA": {Opcode: INCA, Operands: 0},
	"BIOS": {Opcode: BIOS, Operands: 2},
	"CMPA": {Opcode: CMPA, Operands: 1},
	"JEQ":  {Opcode: JEQ, Operands: 1},
	"JUMP": {Opcode: JUMP, Operands: 1},
	"CALL": {Opcode: CALL, Operands: 1},
	"RETN": {Opcode: RETN, Operands: 0},
	"INCI": {Opcode: INCI, Operands: 0},
	"CMPI": {Opcode: CMPI, Operands: 1},
	"SETI": {Opcode: SETI, Operands: 1},
}

type Word uint64

type GMachine struct {
	A, N, P, I     Word
	FlagZ          bool
	Memory         []Word
	Stdout, Stderr io.Writer
}

func New() *GMachine {
	return &GMachine{
		Memory: make([]Word, DefaultMemSize),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (g *GMachine) Run() {
	for {
		opcode := g.Memory[g.P]
		g.P++
		switch opcode {
		case NOOP:
		case HALT:
			return
		case INCA:
			g.A++
		case DECA:
			g.A--
		case SETA:
			g.A = g.Next()
		case SETAM:
			g.A = g.Memory[g.I]
		case BIOS:
			operation := g.Next()
			fileDescriptor := g.Next()
			if operation == IOWrite {
				if fileDescriptor == PortStdout {
					fmt.Fprintf(g.Stdout, "%c", g.A)
					continue
				}
				fmt.Fprintf(g.Stderr, "%c", g.A)
			}
		case CMPA:
			value := g.Next()
			g.FlagZ = g.A == value
		case CMPI:
			value := g.Next()
			g.FlagZ = g.I == value
		case JEQ:
			if !g.FlagZ {
				g.P = g.Memory[g.P]
				continue
			}
			g.P++
		case JUMP:
			g.P = g.Next()
		case CALL:
			g.N = g.P + 1
			g.P = g.Next()
		case RETN:
			g.P = g.N
			g.N = 0
		case INCI:
			g.I++
		}

	}
}

func (g *GMachine) Next() Word {
	next := g.Memory[g.P]
	g.P++
	return next
}

func (g *GMachine) RunProgram(instructions []Word) {
	for i := range instructions {
		g.Memory[i] = instructions[i]
	}
	g.Run()
}

func (g *GMachine) ExecuteBinary(binPath string) error {
	binFile, err := os.Open(binPath)
	if err != nil {
		return err
	}
	defer binFile.Close()
	return g.RunProgramFromReader(binFile)
}

func Assemble(code []string) ([]Word, error) {
	words := []Word{}
	constants := PredefinedConstants
	for pos := 0; pos < len(code); pos++ {
		token := code[pos]
		var word Word

		if strings.HasSuffix(token, ":") {
			fmt.Println("start routine")
		}

		instruction, ok := TranslateTable[token]
		if ok {
			word = instruction.Opcode
		} else {
			value, err := strconv.Atoi(token)
			if err != nil {
				return nil, fmt.Errorf("invalid instruction %q at postion %d", token, pos)
			}
			word = Word(value)

		}
		if instruction.Operands > 0 {
			if pos+instruction.Operands >= len(code) {
				return nil, errors.New("missing operand")
			}
			for count := 0; count < instruction.Operands; count++ {
				// TODO check for square bracket, indicating SETAM operand
				operand := code[pos+1]
				operandWord, ok := constants[operand]
				if !ok {
					temp, err := strconv.Atoi(operand)
					if err != nil {
						return nil, err
					}
					operandWord = Word(temp)
				}
				words = append(words, operandWord)
				pos++
			}
		}
	}
	return words, nil
}

func AssembleFromFile(path string) ([]Word, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	code := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code = append(code, strings.ToUpper(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	words, err := Assemble(code)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func AssembleFromText(text string) ([]Word, error) {
	code := []string{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "#"):
			continue
		}
		for _, item := range strings.Fields(line) {
			code = append(code, strings.ToUpper(item))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	words, err := Assemble(code)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func AssembleFromFileToBinary(inPath, outPath string) error {
	data, err := AssembleFromFile(inPath)
	if err != nil {
		return err
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return WriteWords(outFile, data)
}

func (g *GMachine) RunProgramFromReader(r io.Reader) error {
	words, err := ReadWords(r)
	if err != nil {
		return err
	}
	g.RunProgram(words)

	return nil
}

func ReadWords(r io.Reader) ([]Word, error) {
	rawBytes := make([]byte, 8)
	words := []Word{}
	for {
		_, err := r.Read(rawBytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		b := binary.BigEndian.Uint64(rawBytes)
		words = append(words, Word(b))
	}
	return words, nil
}

func WriteWords(w io.Writer, data []Word) error {
	for _, word := range data {
		rawBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(rawBytes, uint64(word))
		w.Write(rawBytes)
	}
	return nil
}

func RunCLI(path string) error {
	g := New()
	return g.ExecuteBinary(path)
}
