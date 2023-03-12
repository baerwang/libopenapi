// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

func TestRequestBody_MarshalYAML(t *testing.T) {

    req := &RequestBody{
        Description: "beer",
        Required:    true,
        Extensions:  map[string]interface{}{"x-high-gravity": "why not?"},
    }

    rend, _ := req.Render()

    desired := `description: beer
required: true
x-high-gravity: why not?`

    assert.Equal(t, desired, strings.TrimSpace(string(rend)))

}