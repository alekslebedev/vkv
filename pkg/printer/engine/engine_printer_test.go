package engine

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintEngines(t *testing.T) {
	testCases := []struct {
		name     string
		ns       map[string][]string
		opts     []Option
		expected string
		err      bool
	}{
		{
			name: "root & 2 engines",
			ns: map[string][]string{
				"": {"a", "b"},
			},
			opts: []Option{
				ToFormat(Base),
			},
			expected: `a
b
`,
		},
		{
			name: "empty",
			ns: map[string][]string{
				"": {},
			},
			opts: []Option{
				ToFormat(Base),
			},
			expected: ``,
		},
		{
			name: "multi leveled",
			ns: map[string][]string{
				"":     {"a", "b"},
				"a":    {"a1", "a2"},
				"a/a1": {"1", "2"},
			},
			opts: []Option{
				ToFormat(Base),
			},
			expected: `1
2
a
a1
a2
b
`,
		},
		{
			name: "regex",
			ns: map[string][]string{
				"":     {"a", "b"},
				"a":    {"a1", "a2"},
				"a/a1": {"1", "2"},
			},
			opts: []Option{
				ToFormat(Base),
				WithRegex("1"),
			},
			expected: `1
a1
`,
		},
		{
			name: "json",
			ns: map[string][]string{
				"": {"a", "b"},
			},
			opts: []Option{
				ToFormat(JSON),
			},
			expected: `{
  "engines": [
    "a",
    "b"
  ]
}
`,
		},
		{
			name: "yaml",
			ns: map[string][]string{
				"": {"a", "b"},
			},
			opts: []Option{
				ToFormat(YAML),
			},
			expected: `engines:
- a
- b

`,
		},
		{
			name: "yaml",
			ns: map[string][]string{
				"p": {"a", "b"},
			},
			opts: []Option{
				WithNSPrefix(true),
			},
			expected: `p/a
p/b
`,
		},
		{
			name: "invalid regex",
			ns: map[string][]string{
				"p": {"a", "b"},
			},
			opts: []Option{
				WithNSPrefix(true),
				WithRegex("*"),
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		var b bytes.Buffer

		tc.opts = append(tc.opts, WithWriter(&b))

		p := NewPrinter(tc.opts...)

		err := p.Out(tc.ns)

		if tc.err {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tc.expected, b.String(), tc.name)
		}
	}
}
