package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdArgsValidate(t *testing.T) {
	tests := []struct {
		desc    string
		args    CmdArgs
		wantErr bool
	}{
		{
			desc: "ok",
			args: CmdArgs{
				Level: "normal",
				Args:  []string{"LICENSE"},
			},
			wantErr: false,
		},
		{
			desc: "ok: level is hard",
			args: CmdArgs{
				Level: "hard",
				Args:  []string{"LICENSE"},
			},
			wantErr: false,
		},
		{
			desc: "ng: arguments count is 0",
			args: CmdArgs{
				Level: "normal",
				Args:  []string{},
			},
			wantErr: true,
		},
		{
			desc: "ng: illegal level string",
			args: CmdArgs{
				Level: "aiueo",
				Args:  []string{"LICENSE"},
			},
			wantErr: true,
		},
		{
			desc: "ng: file doesn't exist",
			args: CmdArgs{
				Level: "aiueo",
				Args:  []string{"LICENSE", "sushi.txt"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.args.Validate()
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
		})
	}
}
