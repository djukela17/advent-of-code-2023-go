package puzzles

import (
	"os"
	"testing"
)

func TestTrebuchetPuzzle_SolvePartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input",
			args: args{
				input: MustReadFileToString(t, "inputs/example.day_01_part_1.txt"),
			},
			want: 142,
		},
		{
			name: "case with real input that will fail",
			args: args{
				input: MustReadFileToString(t, "inputs/day_01_part_1.txt"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &TrebuchetPuzzle{}
			if got := p.SolvePartOne(tt.args.input); got != tt.want {
				t.Errorf("SolvePartOne() failed!\ngot = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func MustReadFileToString(t *testing.T, filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("could not read file: %v\n", err)
	}

	return string(contents)
}

func TestTrebuchetPuzzle_SolvePartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input",
			args: args{
				input: MustReadFileToString(t, "inputs/example.day_01_part_2.txt"),
			},
			want: 281,
		},
		{
			name: "real input case that will fail",
			args: args{
				input: MustReadFileToString(t, "inputs/day_01_part_2.txt"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &TrebuchetPuzzle{}
			if got := p.SolvePartTwo(tt.args.input); got != tt.want {
				t.Errorf("SolvePartTwo()\n got = %v\nwant %v", got, tt.want)
			}
		})
	}
}
