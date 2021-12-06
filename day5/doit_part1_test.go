package main

import (
	"adventcodingchallenge_2021/utility"
	"reflect"
	"testing"
)

func TestSolutionPart1WithSampleDataFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "sample input",
			args: args{
				fileName: "input_sample.txt",
			},
			want: 5,
		},
		{
			name: "real input",
			args: args{
				fileName: "input.txt",
			},
			want: 7674,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData, err := utility.ParseInputFileIntoStringRows(tt.args.fileName)
			if err != nil {
				t.Error(err)
			}
			if _, got := solution_part_a(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinateRange_determineCoordinatesInRange(t *testing.T) {
	type fields struct {
		start *Coordinate
		end   *Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Coordinate
	}{
		{
			name: "x",
			fields: fields{
				start: &Coordinate{
					x: 0,
					y: 1,
				},
				end: &Coordinate{
					x: 0,
					y: 3,
				},
			},
			want: []*Coordinate{
				&Coordinate{
					x: 0,
					y: 1,
				},
				&Coordinate{
					x: 0,
					y: 2,
				},
				&Coordinate{
					x: 0,
					y: 3,
				},
			},
		},
		{
			name: "x reverse",
			fields: fields{
				start: &Coordinate{
					x: 0,
					y: 3,
				},
				end: &Coordinate{
					x: 0,
					y: 1,
				},
			},
			want: []*Coordinate{
				&Coordinate{
					x: 0,
					y: 1,
				},
				&Coordinate{
					x: 0,
					y: 2,
				},
				&Coordinate{
					x: 0,
					y: 3,
				},
			},
		},

		{
			name: "diag a",
			fields: fields{
				start: &Coordinate{
					x: 1,
					y: 1,
				},
				end: &Coordinate{
					x: 3,
					y: 3,
				},
			},
			want: []*Coordinate{
				&Coordinate{
					x: 1,
					y: 1,
				},
				&Coordinate{
					x: 1,
					y: 2,
				},
				&Coordinate{
					x: 3,
					y: 3,
				},
			},
		},

		{
			name: "diag b",
			fields: fields{
				start: &Coordinate{
					x: 9,
					y: 7,
				},
				end: &Coordinate{
					x: 7,
					y: 9,
				},
			},
			want: []*Coordinate{
				&Coordinate{
					x: 9,
					y: 7,
				},
				&Coordinate{
					x: 8,
					y: 8,
				},
				&Coordinate{
					x: 7,
					y: 9,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoordinateRange{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := c.determineCoordinatesInRange(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineCoordinatesInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
