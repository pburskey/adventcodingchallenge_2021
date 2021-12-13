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
		//{
		//	name: "sample input",
		//	args: args{
		//		fileName: "input_sample2.txt",
		//	},
		//	want: 1656,
		//},
		{
			name: "sample input",
			args: args{
				fileName: "input_sample.txt",
			},
			want: 1656,
		},
		//{
		//	name: "real input",
		//	args: args{
		//		fileName: "input.txt",
		//	},
		//	want: 240123,
		//},
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
