package types

import (
	"errors"
)

type BootSequence struct {
	instructions []Instruction
	accumulator int
	at int
	visits []int
}

func NewBootSequence() BootSequence {
	return BootSequence{
		[]Instruction{},
		0,
		0,
		[]int{},
	}
}

func (b *BootSequence) AddInstruction(instruction Instruction) {
	b.instructions = append(b.instructions, instruction)
	b.visits = append(b.visits, 0)
}

func (b *BootSequence) Run() (int, error) {

	b.Reset()
	done := false

	for !done {
		instruction := b.instructions[b.at]

		if b.visits[b.at] > 0 {
			return b.accumulator, errors.New("Instruction visited twice")
		}

		b.ExecuteInstruction(instruction)

		if b.at >= len(b.instructions) {
			done = true
		}
	}

	return b.accumulator, nil

}

func (b *BootSequence) Reset() () {
	b.accumulator = 0
	b.at = 0
	b.visits = make([]int, len(b.visits))
}

func (b *BootSequence) Repair() () {
	_, err  := b.Run()

	repairIndex := 0

	for err != nil {

		currentInstruction := b.instructions[repairIndex]

		newInstruction := ""
		if currentInstruction.Operation == "nop" {
			newInstruction = "jmp"
		} else if currentInstruction.Operation == "jmp" {
			newInstruction = "nop"
		} else if currentInstruction.Operation == "acc" {
			repairIndex++
			continue
		}

		b.instructions[repairIndex] = NewInstruction(newInstruction, currentInstruction.Arg)

		repairIndex++
		_, err  = b.Run()

		// Set old instruction back if still corrupt
		if err != nil {
			b.instructions[repairIndex - 1] = currentInstruction
		}
	}
}


func (b *BootSequence) ExecuteInstruction(instruction Instruction) {

	b.visits[b.at]++

	switch (instruction.Operation) {
	case "nop":
		b.at++
	case "acc":
		b.accumulator += instruction.Arg
		b.at++
	case "jmp":
		b.at += instruction.Arg
	}
}

type Instruction struct {
	Operation string
	Arg int
}

func NewInstruction(instruction string, arg int) Instruction {
	return Instruction{
		instruction,
		arg,
	}
}
