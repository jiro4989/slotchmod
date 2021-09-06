package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	Level string
	Args  []string
}

func ParseArgs() (*CmdArgs, error) {
	opts := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.StringVar(&opts.Level, "level", "normal", "slot difficulty. [easy|normal|hard]")
	flag.Parse()
	opts.Args = flag.Args()

	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &opts, nil
}

func flagHelpMessage() {
	cmd := os.Args[0]
	fmt.Fprintln(os.Stderr, fmt.Sprintf("%s changes file permissions with a slot", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS] [files...]", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s sample.txt", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")

	flag.PrintDefaults()
}

func (c *CmdArgs) Validate() error {
	if len(c.Args) < 1 {
		return fmt.Errorf("Must need files")
	}

	_, ok := slotIntervalTime[c.Level]
	if !ok {
		return fmt.Errorf("-level must be 'eash' or 'normal' or 'hard'.")
	}

	for _, file := range c.Args {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s file doesn't exist.", file)
		}
	}

	return nil
}
