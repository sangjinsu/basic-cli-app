package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("## TestMain Start")
	exitVal := m.Run()
	fmt.Println("## TestMain End")
	os.Exit(exitVal)
}

func Test_parseArgs(t *testing.T) {
	type args struct {
		writer *bufio.Writer
		args   []string
	}
	tests := []struct {
		name    string
		args    args
		want    config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "help",
			args: args{
				writer: bufio.NewWriter(os.Stdin),
				args:   []string{"-h"},
			},
			want: config{
				numTimes: 0,
			},
			wantErr: true,
		},
		{
			name: "normal case",
			args: args{
				writer: bufio.NewWriter(os.Stdin),
				args:   []string{"-n", "10"},
			},
			want: config{
				numTimes: 10,
			},
			wantErr: false,
		},
		{
			name: "not number",
			args: args{
				writer: bufio.NewWriter(os.Stdin),
				args:   []string{"-n", "abc"},
			},
			want: config{
				numTimes: 0,
			},
			wantErr: true,
		},
		{
			name: "invalid argument",
			args: args{
				writer: bufio.NewWriter(os.Stdin),
				args:   []string{"-n", "1", "invalid", "argument"},
			},
			want: config{
				numTimes: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args.writer, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
