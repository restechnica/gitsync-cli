package target

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_SelectTarget(t *testing.T) {
	var availableTargets = []Target{
		NewGitTarget(),
	}

	type Test struct {
		Target string
		Name   string
		Want   Target
	}

	var tests = []Test{
		{Name: "SelectTargetSuccessfully", Target: "blabla.git", Want: NewGitTarget()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got, err = SelectTarget(test.Target, availableTargets)

			assert.NoError(t, err)
			assert.IsType(t, test.Want, got, `want: "%s, got: "%s"`, test.Want, got)
		})
	}

	type ErrorTest struct {
		AvailableTargets []Target
		Name             string
		Target           string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnUnavailableTarget", Target: "git", AvailableTargets: []Target{}},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var _, got = SelectTarget(test.Target, test.AvailableTargets)
			assert.Error(t, got)
		})
	}
}
