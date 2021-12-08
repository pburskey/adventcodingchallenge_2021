package main

import (
	"strings"
	"testing"
)

func Test_displaySegment_decode(t *testing.T) {

	tests := []struct {
		name       string
		cypherKeys string
		cypher     string
		want       int
		wantErr    bool
	}{
		//{
		//	name:       "",
		//	cypherKeys: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab",
		//	cypher:     "cdfeb",
		//	want:       5,
		//	wantErr:    false,
		//},
		//{
		//	name:       "",
		//	cypherKeys: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab",
		//	cypher:     "fcadb",
		//	want:       3,
		//	wantErr:    false,
		//},
		{
			name:       "",
			cypherKeys: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab",
			cypher:     "cdfeb",
			want:       5,
			wantErr:    false,
		},
		//{
		//	name:       "",
		//	cypherKeys: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab",
		//	cypher:     "cdbaf",
		//	want:       3,
		//	wantErr:    false,
		//},
		//{
		//	name:       "",
		//	cypherKeys: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
		//	cypher:     "fdgacbe",
		//	want:       8,
		//	wantErr:    false,
		//},
		{
			name:       "",
			cypherKeys: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
			cypher:     "cefdb",
			want:       3,
			wantErr:    false,
		},
		//{
		//	name:       "",
		//	cypherKeys: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
		//	cypher:     "cefbgd",
		//	want:       9,
		//	wantErr:    false,
		//},
		//{
		//	name:       "",
		//	cypherKeys: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
		//	cypher:     "gcbe",
		//	want:       4,
		//	wantErr:    false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &displaySegment{
				cypher: tt.cypher,
			}

			keys := &cypherKeys{
				public: strings.Split(tt.cypherKeys, " "),
			}
			keys.loadKeys()

			if !s.decode(keys) {
				t.Errorf("decode() error = %v", tt.wantErr)
				return
			}
			if s.value != tt.want {
				t.Errorf("decode(%s) got = %v, want %v", tt.cypher, s.value, tt.want)
			}
		})
	}
}
