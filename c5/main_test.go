package main

import (
	"testing"
)

func Test_calculate(t *testing.T) {
	type args struct {
		a  int
		b  int
		op string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addition with positive outcome",
			args: args{
				a:  2,
				b:  2,
				op: "add",
			},
			want: 4,
		},
		{
			name: "Addition with negative outcome",
			args: args{
				a:  23,
				b:  9,
				op: "add",
			},
			want: 4,
			// want: 32,
		},
		{
			name: "Substraction with positive outcome",
			args: args{
				a:  23,
				b:  9,
				op: "substract",
			},
			want: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.a, tt.args.b, tt.args.op); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vocals(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Text one true",
			args: args{
				text: "Bojan",
			},
			want: true,
		},
		{
			name: "Text two true",
			args: args{
				text: "Alexandar",
			},
			want: true,
		},
		{
			name: "Text two false",
			args: args{
				text: "Ivana",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vocals(tt.args.text); got != tt.want {
				t.Errorf("vocals() = %v, want %v", got, tt.want)
			}
		})
	}
}
