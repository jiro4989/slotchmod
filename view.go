package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

var (
	bigNumbers = [][]string{
		{
			`    _    `,
			`   / |   `,
			`   | |   `,
			`   | |   `,
			`   |_|   `,
			`         `,
		},
		{
			`  ____   `,
			` |___ \  `,
			`   __) | `,
			`  / __/  `,
			` |_____| `,
			`         `,
		},
		{
			`  _____  `,
			` |___ /  `,
			`   |_ \  `,
			`  ___) | `,
			` |____/  `,
			`         `,
		},
		{
			` _  _    `,
			`| || |   `,
			`| || |_  `,
			`|__   _| `,
			`   |_|   `,
			`         `,
		},
		{
			`  ____   `,
			` | ___|  `,
			` |___ \  `,
			`  ___) | `,
			` |____/  `,
			`         `,
		},
		{
			`   __    `,
			`  / /_   `,
			` | '_ \  `,
			` | (_) | `,
			`  \___/  `,
			`         `,
		},
		{
			`  _____  `,
			` |___  | `,
			`    / /  `,
			`   / /   `,
			`  /_/    `,
			`         `,
		},
		{
			`   ___   `,
			`  ( _ )  `,
			`  / _ \  `,
			` | (_) | `,
			`  \___/  `,
			`         `,
		},
		{
			`   ___   `,
			`  / _ \  `,
			` | (_) | `,
			`  \__, | `,
			`    /_/  `,
			`         `,
		},
	}
)

func DrawSlot(s *Slot) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	idx := s.CurrentSlotIndex()
	pv := s.PreviousValue()
	nv := s.NextValue()
	slots := s.Slots()

	p := [3]string{" ", " ", " "}
	p[idx] = fmt.Sprintf("%d", pv)

	n := [3]string{" ", " ", " "}
	n[idx] = fmt.Sprintf("%d", nv)

	rows := []string{
		fmt.Sprintf("      %s %s %s", p[0], p[1], p[2]),
		fmt.Sprintf("chmod %d %d %d", slots[0], slots[1], slots[2]),
		fmt.Sprintf("      %s %s %s", n[0], n[1], n[2]),
	}

	for y, row := range rows {
		for x, r := range row {
			termbox.SetChar(x, y, r)
		}
	}

	termbox.Flush()
}
