// Copyright 2024 "6543". All rights reserved.
// SPDX-License-Identifier: MIT

package optional_test

import (
	"testing"

	"github.com/6543/go-optional/v2"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestOptionalToYaml(t *testing.T) {
	tests := []struct {
		name string
		obj  *testSerializationStruct
		want string
	}{
		{
			name: "empty",
			obj:  new(testSerializationStruct),
			want: `normal_string: ""
normal_bool: false
optional_two_bool: null
optional_two_string: null
`,
		},
		{
			name: "some",
			obj: &testSerializationStruct{
				NormalString: "a string",
				NormalBool:   true,
				OptBool:      optional.Some(false),
				OptString:    optional.Some(""),
			},
			want: `normal_string: a string
normal_bool: true
optional_bool: false
optional_string: ""
optional_two_bool: null
optional_two_string: null
`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b, err := yaml.Marshal(tc.obj)
			assert.NoError(t, err)
			assert.EqualValues(t, tc.want, string(b), "yaml module returned unexpected")
		})
	}
}

func TestOptionalFromYaml(t *testing.T) {
	tests := []struct {
		name string
		data string
		want testSerializationStruct
	}{
		{
			name: "empty",
			data: ``,
			want: testSerializationStruct{},
		},
		{
			name: "empty but init",
			data: `normal_string: ""
normal_bool: false
optional_bool:
optional_two_bool:
optional_two_string:
`,
			want: testSerializationStruct{},
		},
		{
			name: "some",
			data: `
normal_string: a string
normal_bool: true
optional_bool: false
optional_string: ""
optional_two_bool: null
optional_twostring: null
`,
			want: testSerializationStruct{
				NormalString: "a string",
				NormalBool:   true,
				OptBool:      optional.Some(false),
				OptString:    optional.Some(""),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var obj testSerializationStruct
			err := yaml.Unmarshal([]byte(tc.data), &obj)
			assert.NoError(t, err)
			assert.EqualValues(t, tc.want, obj, "yaml module returned unexpected")
		})
	}
}
