// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package shields

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		label   string
		message string
	}{
		{"foo", "bar"},
		{"fizz", "buzz"},
	}
	for _, test := range cases {
		badge := New(test.label, test.message)
		assert.Equal(t, 1, badge.Version)
		assert.Equal(t, test.label, badge.Label)
		assert.Equal(t, test.message, badge.Message)
	}
}

func TestFromRequest(t *testing.T) {
	label := "foo"
	message := "bar"
	values := url.Values{}
	values.Set("schemaVersion", "2") // ignored
	values.Set("label", "fizz")
	values.Set("labelColor", ColorGreen)
	values.Set("message", "buzz") // ignored
	values.Set("color", ColorBlue)
	values.Set("isError", "1") // ignored
	values.Set("namedLogo", "go")
	values.Set("logoSvg", "logoSvg")
	values.Set("logoColor", "white")
	values.Set("logoWidth", "40")
	values.Set("logoPosition", "left")
	values.Set("style", StylePlastic)
	values.Set("cacheSeconds", "500") // ignored
	req := httptest.NewRequest(http.MethodGet, "/?"+values.Encode(), nil)
	badge := New(label, message)
	err := badge.ParseRequest(req)
	assert.Nil(t, err)
	assert.Equal(t, "fizz", badge.Label)
	assert.Equal(t, ColorGreen, badge.LabelColor)
	assert.Equal(t, message, badge.Message)
	assert.Equal(t, 1, badge.Version)
	assert.Equal(t, ColorBlue, badge.Color)
	assert.Equal(t, false, badge.IsError)
	assert.Equal(t, "go", badge.NamedLogo)
	assert.Equal(t, "white", badge.LogoColor)
	assert.Equal(t, "40", badge.LogoWidth)
	assert.Equal(t, "left", badge.LogoPosition)
	assert.Equal(t, "logoSvg", badge.LogoSvg)
	assert.Equal(t, StylePlastic, badge.Style)
	assert.Equal(t, 0, badge.CacheSeconds)
}
