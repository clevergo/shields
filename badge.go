// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package shields

import (
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func init() {
	decoder.IgnoreUnknownKeys(true)
}

// Badge is the endpoint schema of shields.io.
type Badge struct {
	// Schema version. Always the number 1.
	// Required.
	Version int `json:"schemaVersion" schema:"-"`

	// The left text, or the empty string to omit the left side of the badge.
	// Required.
	Label string `json:"label" schema:"label"`

	// The right text.
	// Required.
	Message string `json:"message" schema:"-"`

	// The right color. Supports the eight named colors above,
	// as well as hex, rgb, rgba, hsl, hsla and css named colors.
	// Default: lightgrey.
	Color string `json:"color,omitempty" schema:"color"`

	// The left color.
	// Default: grey.
	LabelColor string `json:"labelColor,omitempty" schema:"labelColor"`

	// true to treat this as an error badge. This prevents the user from overriding the color.
	// In the future it may affect cache behavior.
	// Default: false.
	IsError bool `json:"isError,omitempty" schema:"-"`

	// One of the named logos supported by Shields or simple-icons(https://simpleicons.org/).
	// Default: none.
	NamedLogo string `json:"namedLogo,omitempty" schema:"namedLogo"`

	// An SVG string containing a custom logo.
	// Default: none.
	LogoSvg string `json:"logoSvg,omitempty" schema:"logoSvg"`

	// Logo color.
	// Default: none.
	LogoColor string `json:"logoColor,omitempty" schema:"logoColor"`

	// The horizontal space to give to the logo.
	// Default: none.
	LogoWidth string `json:"logoWidth,omitempty" schema:"logoWidth"`

	// Logo position.
	// Default: none.
	LogoPosition string `json:"logoPosition,omitempty" schema:"logoPosition"`

	// The default template to use.
	// Default: flat.
	Style string `json:"style,omitempty" schema:"style"`

	// Set the HTTP cache lifetime in seconds, which should be respected by the Shields' CDN and downstream users.
	// Default: 300.
	CacheSeconds int `json:"cacheSeconds,omitempty" schema:"-"`
}

// New returns a badge with the given label and message.
func New(label, message string) *Badge {
	return &Badge{
		Version: 1,
		Label:   label,
		Message: message,
	}
}

// ParseRequest parses request query params.
func (badge *Badge) ParseRequest(req *http.Request) error {
	return decoder.Decode(badge, req.URL.Query())
}
