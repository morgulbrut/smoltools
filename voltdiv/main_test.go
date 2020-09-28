package main

import (
	"testing"
)

func Test_parseResistor(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseResistor(tt.args.r); got != tt.want {
				t.Errorf("parseResistor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcH(t *testing.T) {
	type args struct {
		r1 float64
		r2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcH(tt.args.r1, tt.args.r2); got != tt.want {
				t.Errorf("calcH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
