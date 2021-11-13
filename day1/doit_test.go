package main

import (
	"advent2021/utility"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSolutionWithSampleData(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Sunny Day Easy data 1",
			args: args{
				data: []string{
					"1000",
					"2000",
				},
			},
			want: 3000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionWithSampleDataFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Sunny Day 1",
			args: args{
				fileName: "data.txt",
			},
			want: 4695,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData, err := utility.ParseInputFileIntoRows(filepath.Join(utility.AssembleFilePathToDay(day), tt.args.fileName))
			if err != nil {
				t.Error(err)
			}
			if got := Solution(inputData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
