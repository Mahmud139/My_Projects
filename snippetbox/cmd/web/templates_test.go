package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2022, 07, 29, 07, 57, 0, 0, time.Local)
	hd := humanDate(tm)

	if hd != "29 Jul 2022 at 07:57" {
		t.Errorf("want %q; got %q", "29 Jul 2022 at 07:57", hd)
	}
}