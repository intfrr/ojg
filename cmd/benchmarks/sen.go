// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/sen"
)

func senParse(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		log.Fatal(err)
	}
	b.ResetTimer()
	p := &sen.Parser{}
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample); err != nil {
			log.Fatal(err)
		}
	}
}

func senParseReader(b *testing.B) {
	var p sen.Parser
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err = p.ParseReader(f); err != nil {
			log.Fatal(err)
		}
	}
}

func senString(b *testing.B) {
	data := loadSample()
	opt := sen.Options{OmitNil: true}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = sen.String(data, &opt)
	}
}

func senStringIndent(b *testing.B) {
	data := loadSample()
	b.ResetTimer()
	opt := sen.Options{OmitNil: true, Indent: 2}
	for n := 0; n < b.N; n++ {
		_ = sen.String(data, &opt)
	}
}

func senStringSort(b *testing.B) {
	data := loadSample()
	opt := sen.Options{OmitNil: true, Indent: 2, Sort: true}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = sen.String(data, &opt)
	}
}
