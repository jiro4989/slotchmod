package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	slotIntervalTime = map[string]int{
		"easy":   200,
		"normal": 100,
		"hard":   50,
	}
)

func main() {
	args, err := ParseArgs()
	if err != nil {
		Err(err)
		os.Exit(1)
	}

	interval := slotIntervalTime[args.Level]
	slot := NewSlot(0, interval)

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	go clock(slot)
	waitKeyInput(slot)
	termbox.Close()

	changeMode(slot, args.Args)
}

func clock(s *Slot) {
	for !s.IsFinished() {
		s.Switch()
		drawSlot(s)
		time.Sleep(time.Duration(s.IntervalTime()) * time.Millisecond)
	}
}

func waitKeyInput(s *Slot) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD:
				return
			case termbox.KeyEnter:
				s.Select()
			}
			switch ev.Ch {
			case 'q':
				return
			}
		}
		if s.IsFinished() {
			return
		}
	}
}

func changeMode(s *Slot, files []string) {
	slots := s.Slots()
	for _, file := range files {
		perm := fmt.Sprintf("0%d%d%d", slots[0], slots[1], slots[2])
		t := fmt.Sprintf("chmod %s %s", perm, file)
		fmt.Println(t)

		mode, err := strconv.ParseUint(perm, 8, 32)
		if err != nil {
			panic(err)
		}
		if err := os.Chmod(file, os.FileMode(mode)); err != nil {
			panic(err)
		}
	}
	if slots[0] == slots[1] && slots[1] == slots[2] {
		fmt.Println("BINGO🎉")
	}
}
