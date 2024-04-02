package target

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalGitTarget_GetName(t *testing.T) {
	t.Run("CheckName", func(t *testing.T) {
		var want = LocalGit
		var got = LocalGitTarget{}.GetName()

		assert.Equal(t, want, got, `want: '%s', got: '%s'`, want, got)
	})
}

func TestLocalGitTarget_IsCompatible(t *testing.T) {
	type Test struct {
		Id   string
		Name string
		Want bool
	}

	var tests = []Test{
		{Name: "RelativePathIsCompatible", Id: "some/relative/path", Want: true},
		{Name: "RelativeHomePathIsCompatible", Id: "~/some/relative/path", Want: true},
		{Name: "AbsolutePathIsCompatible", Id: "/some/absolute/path", Want: true},
		{Name: "HttpIsNotCompatible", Id: "http://github.com/restechnica/gitsync-cli.git", Want: false},
		{Name: "HttpsIsNotCompatible", Id: "https://github.com/restechnica/gitsync-cli.git", Want: false},
		{Name: "SshIsNotCompatible", Id: "git@github.com:restechnica/gitsync-cli.git", Want: false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var target = NewLocalGitTarget()
			var got = target.IsCompatible(test.Id)

			assert.IsType(t, test.Want, got, `want: '%s, got: '%s'`, test.Want, got)
		})
	}
}
