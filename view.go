package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

type DrawStyle int

const (
	DrawStyleSimple DrawStyle = iota
	DrawStyleBig
)

var (
	bigEmptyChmod = [6]string{
		`                                  `,
		`                                  `,
		`                                  `,
		`                                  `,
		`                                  `,
		`                                  `,
	}
	bigChmod = [6]string{
		`      _                         _ `,
		`  ___| |__  _ __ ___   ___   __| |`,
		" / __| '_ \\| '_ ` _ \\ / _ \\ / _` |",
		`| (__| | | | | | | | | (_) | (_| |`,
		` \___|_| |_|_| |_| |_|\___/ \__,_|`,
		`                                  `,
	}
	emptyNumber = [6]string{
		`         `,
		`         `,
		`         `,
		`         `,
		`         `,
		`         `,
	}
	bigNumbers = [10][6]string{
		{
			`   ___   `,
			`  / _ \  `,
			` | | | | `,
			` | |_| | `,
			`  \___/  `,
			`         `,
		},
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

func DrawSlot(s *Slot, st DrawStyle) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	idx := s.CurrentSlotIndex()
	pv := s.PreviousValue()
	nv := s.NextValue()
	slots := s.Slots()

	switch st {
	case DrawStyleSimple:
		drawSimple(slots, idx, pv, nv)
	case DrawStyleBig:
		drawBig(slots, idx, pv, nv)
	default:
		drawSimple(slots, idx, pv, nv)
	}

	termbox.Flush()
}

func drawSimple(slots [3]int, idx, pv, nv int) {
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
}

func drawBig(slots [3]int, idx, pv, nv int) {
	p := [3][6]string{emptyNumber, emptyNumber, emptyNumber}
	p[idx] = bigNumbers[pv]

	s := [3][6]string{
		bigNumbers[slots[0]],
		bigNumbers[slots[1]],
		bigNumbers[slots[2]],
	}

	n := [3][6]string{emptyNumber, emptyNumber, emptyNumber}
	n[idx] = bigNumbers[nv]

	genRow := func(arr [3][6]string, pre [6]string) []string {
		max := len(arr[0])
		var result []string
		for i := 0; i < max; i++ {
			row := fmt.Sprintf("%s %s %s %s", pre[i], arr[0][i], arr[1][i], arr[2][i])
			result = append(result, row)
		}
		return result
	}

	var rows []string
	rows = append(rows, genRow(p, bigEmptyChmod)...)
	rows = append(rows, genRow(s, bigChmod)...)
	rows = append(rows, genRow(n, bigEmptyChmod)...)

	for y, row := range rows {
		for x, r := range row {
			termbox.SetChar(x, y, r)
		}
	}
}
