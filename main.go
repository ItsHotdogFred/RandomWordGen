package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	words = []string{"apple", "banana", "cherry", "date", "elder", "fig",
		"grape", "honey", "lemon", "mango", "orange", "peach", "pear",
		"plum", "berry", "fruit", "tree", "leaf", "seed", "root", "stem",
		"bloom", "petal", "thorn", "vine", "juice", "sweet", "sour",
		"tart", "ripe", "core", "flesh", "skin", "pit", "nut", "cheese", "milk"}

	amount        = 2
	statePath     = "counter.txt"
	totalPossible = pow(len(words), amount)

	// LCG parameters for obfuscation
	lcgA = uint64(1664525)
	lcgC = uint64(1013904223)
)

// getNextString generates the next unique string of words
func GetNextString() string {
	currentID := loadCounter()

	if currentID >= uint64(totalPossible) {
		return "Add more words"
	}

	combo := idToCombo(currentID)

	// Save the incremented counter for next time
	if err := saveCounter(currentID + 1); err != nil {
		log.Fatal(err)
	}

	return strings.Join(combo, "")
}

func idToCombo(id uint64) []string {
	// Apply LCG to make sequence less predictable
	obfuscatedPos := (lcgA*id + lcgC) % uint64(totalPossible)

	// Convert to word combination
	combo := make([]string, amount)
	pos := int(obfuscatedPos)

	for i := amount - 1; i >= 0; i-- {
		combo[i] = words[pos%len(words)]
		pos /= len(words)
	}

	return combo
}

func loadCounter() uint64 {
	data, err := os.ReadFile(statePath)
	if os.IsNotExist(err) {
		return 0
	}
	if err != nil {
		log.Fatalf("loading counter: %v", err)
	}

	counter, err := strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		log.Fatalf("parsing counter: %v", err)
	}
	return counter
}

func saveCounter(counter uint64) error {
	return os.WriteFile(statePath, []byte(strconv.FormatUint(counter, 10)), 0o644)
}

func pow(base, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res
}
