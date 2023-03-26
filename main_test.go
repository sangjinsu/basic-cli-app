package main

import (
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
		args []string
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
				[]string{"-h"},
			},
			want: config{
				printUsage: true, numTimes: 0,
			},
			wantErr: false,
		},
		{
			name: "normal case",
			args: args{
				[]string{"10"},
			},
			want: config{
				printUsage: false, numTimes: 10,
			},
			wantErr: false,
		},
		{
			name: "not number",
			args: args{
				[]string{"abc"},
			},
			want: config{
				printUsage: false, numTimes: 0,
			},
			wantErr: true,
		},
		{
			name: "invalid argument",
			args: args{
				[]string{"1", "invalid argument"},
			},
			want: config{
				printUsage: false, numTimes: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args.args)
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
