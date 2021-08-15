package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "tern/internal"
)

func TestRun(t *testing.T) {
	t.Run(`shell mode`, func(t *testing.T) {
		t.Run(`truthy`, func(t *testing.T) {
			tests := []test{
				{
					name:     "0",
					expected: "yes",
					given:    []string{"0", "yes", "no"},
				},
				{
					name:     "1",
					expected: "yes",
					given:    []string{"1", "yes", "no"},
				},
				{
					name:     "false",
					expected: "yes",
					given:    []string{"false", "yes", "no"},
				},
				{
					name:     "true",
					expected: "yes",
					given:    []string{"true", "yes", "no"},
				},
			}

			check(t, tests)
		})

		t.Run(`falsy`, func(t *testing.T) {
			tests := []test{
				{
					name:     "empty string",
					expected: "no",
					given:    []string{"", "yes", "no"},
				},
			}

			check(t, tests)
		})
	})

	t.Run(`bool mode`, func(t *testing.T) {
		t.Run(`truthy`, func(t *testing.T) {
			tests := []test{
				{
					name:     "true",
					expected: "yes",
					given:    []string{"-b", "true", "yes", "no"},
				},
				{
					name:     "TRUE",
					expected: "yes",
					given:    []string{"-b", "TRUE", "yes", "no"},
				},
				{
					name:     "1",
					expected: "yes",
					given:    []string{"-b", "1", "yes", "no"},
				},
			}

			check(t, tests)
		})

		t.Run(`falsy`, func(t *testing.T) {
			tests := []test{
				{
					name:     "empty string",
					expected: "no",
					given:    []string{"-b", "", "yes", "no"},
				},
				{
					name:     "0",
					expected: "no",
					given:    []string{"-b", "0", "yes", "no"},
				},
				{
					name:     "false",
					expected: "no",
					given:    []string{"-b", "false", "yes", "no"},
				},
				{
					name:     "FALSE",
					expected: "no",
					given:    []string{"-b", "FALSE", "yes", "no"},
				},
			}

			check(t, tests)
		})
	})

	t.Run(`incorrect invocation`, func(t *testing.T) {
		tests := []struct {
			name     string
			expected error
			given    []string
		}{
			{
				name:     "unsupported mode",
				expected: ErrUnsupportedMode,
				given:    []string{"-f", "foo", "yes", "no"},
			},
			{
				name:     "incorrect number of arguments",
				expected: ErrIncorrectNumberOfArguments,
				given:    []string{"foo", "bar"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := Run(tt.given)
				assert.ErrorIs(t, tt.expected, err)
			})
		}
	})
}

func check(t *testing.T, tests []test) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := Run(tt.given)
			assert.Equal(t, tt.expected, res)
		})
	}
}

type (
	test struct {
		name     string
		expected string
		given    []string
	}
)
