package main

import (
	"github.com/neuronpool/go-brainfuck"
)

// ============================================================
// Types and globals
// ============================================================

type instance struct {
	vm        *brainfuck.VM
	userID    string
	channelID string
}

// ============================================================
// Constructor
// ============================================================

func newInstance(userID, channelID string) *instance {
	return &instance{
		vm:        new(brainfuck.VM),
		userID:    userID,
		channelID: channelID,
	}
}
