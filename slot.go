package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

const (
	slotMinValue = 0
	slotMaxValue = 7
)

type Slot struct {
	slots            [3]int
	currentSlotIndex int
	isFinished       bool
}

func NewSlot(seed int64) *Slot {
	if seed == 0 {
		i, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		seed = i.Int64()
	}
	rand.Seed(seed)

	s := Slot{}
	for i := 0; i < 3; i++ {
		slotValue := rand.Intn(slotMaxValue + 1)
		s.slots[i] = slotValue
	}

	return &s
}

func (s *Slot) Switch() {
	s.slots[s.currentSlotIndex] = s.NextValue()
}

func (s *Slot) Select() {
	if 2 <= s.currentSlotIndex {
		s.isFinished = true
		return
	}
	s.currentSlotIndex++
}

func (s *Slot) IsFinished() bool {
	return s.isFinished
}

func (s *Slot) Slots() [3]int {
	return s.slots
}

func (s *Slot) PreviousValue() int {
	v := s.slots[s.currentSlotIndex] - 1
	if v < slotMinValue {
		v = slotMaxValue
	}
	return v
}

func (s *Slot) CurrentValue() int {
	return s.slots[s.currentSlotIndex]
}

func (s *Slot) CurrentSlotIndex() int {
	return s.currentSlotIndex
}

func (s *Slot) NextValue() int {
	v := s.slots[s.currentSlotIndex] + 1
	if slotMaxValue < v {
		v = slotMinValue
	}
	return v
}
