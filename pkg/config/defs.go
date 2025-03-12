package config

import (
	_ "embed"
)

//go:embed schema.json
var Schema string

//go:embed example.yml
var Example string
