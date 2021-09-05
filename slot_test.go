package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlotNewSlot(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(NewSlot(0))
	assert.NotNil(NewSlot(1))
}

func TestSlotPreviousValue(t *testing.T) {
	tests := []struct {
		desc string
		slot Slot
		want int
	}{
		{
			desc: "0 -> 7",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 0,
			},
			want: 7,
		},
		{
			desc: "6 -> 5",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 1,
			},
			want: 5,
		},
		{
			desc: "7 -> 6",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 2,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.slot.PreviousValue()
			assert.Equal(tt.want, got)
		})
	}
}

func TestSlotCurrentValue(t *testing.T) {
	tests := []struct {
		desc string
		slot Slot
		want int
	}{
		{
			desc: "0",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 0,
			},
			want: 0,
		},
		{
			desc: "6",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 1,
			},
			want: 6,
		},
		{
			desc: "7",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 2,
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.slot.CurrentValue()
			assert.Equal(tt.want, got)
		})
	}
}

func TestSlotNextValue(t *testing.T) {
	tests := []struct {
		desc string
		slot Slot
		want int
	}{
		{
			desc: "0 -> 1",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 0,
			},
			want: 1,
		},
		{
			desc: "6 -> 7",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 1,
			},
			want: 7,
		},
		{
			desc: "7 -> 0",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 2,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := tt.slot.NextValue()
			assert.Equal(tt.want, got)
		})
	}
}

func TestSlotSelect(t *testing.T) {
	tests := []struct {
		desc                 string
		slot                 Slot
		wantIsFinished       bool
		wantCurrentSlotIndex int
	}{
		{
			desc: "turn ON isFinished flag when slot index is 2, and slot index is 2",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 2,
			},
			wantIsFinished:       true,
			wantCurrentSlotIndex: 2,
		},
		{
			desc: "turn OFF isFinished flag when slot index is 0 or 1, and increments slot index",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 0,
			},
			wantIsFinished:       false,
			wantCurrentSlotIndex: 1,
		},
		{
			desc: "turn OFF isFinished flag when slot index is 0 or 1, and increments slot index",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 1,
			},
			wantIsFinished:       false,
			wantCurrentSlotIndex: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			tt.slot.Select()
			assert.Equal(tt.wantIsFinished, tt.slot.isFinished)
			assert.Equal(tt.wantCurrentSlotIndex, tt.slot.currentSlotIndex)
		})
	}
}

func TestSlotSwitch(t *testing.T) {
	tests := []struct {
		desc string
		slot Slot
		want [3]int
	}{
		{
			desc: "increments slots[0]",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 0,
			},
			want: [3]int{1, 6, 7},
		},
		{
			desc: "increments slots[1]",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 1,
			},
			want: [3]int{0, 7, 7},
		},
		{
			desc: "increments slots[2], and set 0",
			slot: Slot{
				slots:            [3]int{0, 6, 7},
				currentSlotIndex: 2,
			},
			want: [3]int{0, 6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			tt.slot.Switch()
			assert.Equal(tt.want, tt.slot.slots)
		})
	}
}
