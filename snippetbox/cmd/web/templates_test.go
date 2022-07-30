package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	/* Unit testing
	tm := time.Date(2022, 07, 29, 07, 57, 0, 0, time.Local)
	hd := humanDate(tm)

	if hd != "29 Jul 2022 at 07:57" {
		t.Errorf("want %q; got %q", "29 Jul 2022 at 07:57", hd)
	} */

	//Table-Driven testing
	// Create a slice of anonymous structs containing the test case name, 
	// input to our humanDate() function (the tm field), and expected output 
	// (the want field).
	tests := []struct {
		name string
		tm time.Time
		want string
	}{
		{
			name: "Local",
			tm: time.Date(2022, 07, 29, 10, 15, 0, 0, time.Local),
			want: "29 Jul 2022 at 10:15",
		},
		{
			name: "Empty",
			tm: time.Time{},
			want: "",
		}, 
		{
			name: "CET",
			tm: time.Date(2022, 07, 29, 10, 30, 0, 0, time.FixedZone("UTC", 6*60*60)),
			want: "29 Jul 2022 at 04:30",
		},
	}

	//loop over the test cases
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The 
		// first parameter to this is the name of the test (which is used to 
		// identify the sub-test in any log output) and the second parameter is 
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}