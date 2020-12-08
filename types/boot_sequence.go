package types

import (
	"errors"
)

type BootSequence struct {
	instructions []Instruction
	accumulator int
	at int
}

func NewBootSequence() BootSequence {
	return BootSequence{
		[]Instruction{},
		0,
		0,
	}
}

func (b *BootSequence) AddInstruction(instruction Instruction) {
	b.instructions = append(b.instructions, instruction)
}

func (b *BootSequence) Run() (int, error) {

	done := false

	for !done {
		instruction := &b.instructions[b.at]

		if instruction.Visits > 0 {
			return b.accumulator, errors.New("Instruction visited twice")
		}

		b.ExecuteInstruction(instruction)
	}

	return b.accumulator, nil

}

func (b *BootSequence) ExecuteInstruction(instruction *Instruction) {

	switch (instruction.Operation) {
	case "nop":
		b.at++
	case "acc":
		b.accumulator += instruction.Arg
		b.at++
	case "jmp":
		b.at += instruction.Arg
	}

	instruction.Visits++
}

type Instruction struct {
	Operation string
	Arg int
	Visits int
}

func NewInstruction(instruction string, arg int) Instruction {
	return Instruction{
		instruction,
		arg,
		0,
	}
}
