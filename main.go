package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	flag.Usage = func() {
		cmd := os.Args[0]
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s changes file permissions with a slot", cmd))
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS] [files...]", cmd))
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s sample.txt", cmd))
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "[ERR] Must need files")
		os.Exit(1)
	}

	slot := NewSlot(0)

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	go clock(slot)
	waitKeyInput(slot)
	termbox.Close()

	changeMode(slot, args)
}

func clock(s *Slot) {
	for !s.IsFinished() {
		s.Switch()
		drawSlot(s)
		time.Sleep(100 * time.Millisecond)
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
