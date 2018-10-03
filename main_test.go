package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"meow"},
		{"fasfas"},
	}

	for tt := range tests {
		t.Run(tests[tt].name, func(t *testing.T) {
			main()
		})
	}
}
