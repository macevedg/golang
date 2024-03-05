package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {

	testcases := map[string]struct {
		in  []Human
		err error
	}{
		"test1": {
			in:  nil,
			err: assert.AnError,
		},
		"test2": {
			in:  []Human{},
			err: assert.AnError,
		},
		"test3": {
			in: []Human{
				&Student{
					Name: "Mauricio",
				},
			},
			err: nil,
		},
	}

	for name, tt := range testcases {

		t.Run(name, func(t *testing.T) {
			printer := Printer{}

			val, err := printer.printNames(tt.in)

			if tt.err != nil {
				assert.Error(t, err)
				return
			}

			assert.NotEmpty(t, val)
		})
	}

}
