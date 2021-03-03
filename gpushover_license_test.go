// Package gpushover is a Go wrapper for Pushover's notification API.
//
// Copyright © 2021 Jeffrey H. Johnson. <trnsz@pobox.com>
// Copyright © 2021 José Manuel Díez. <j.diezlopez@protonmail.ch>
// Copyright © 2021 Gridfinity, LLC. <admin@gridfinity.com>
// Copyright © 2014 Damian Gryski. <damian@gryski.com>
// Copyright © 2014 Adam Lazzarato.
//
// All Rights reserved.
//
// All use of this code is governed by the MIT license.
// The complete license is available in the LICENSE file.

package gpushover_test

import (
	"fmt"
	"testing"

	u "github.com/johnsonjh/leaktestfe"
	licn "go4.org/legal"
)

func TestLicense(
	t *testing.T,
) {
	defer u.Leakplug(
		t,
	)
	licenses := licn.Licenses()
	if len(
		licenses,
	) == 0 {
		t.Fatal(
			"\ngpushover_license_test.TestLicense.licenses FAILURE",
		)
	} else {
		t.Log(
			fmt.Sprintf(
				"\n%v\n",
				licenses,
			),
		)
	}
}
