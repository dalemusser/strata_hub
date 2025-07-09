package config

import "github.com/dalemusser/gowebcore/config"

// Config embeds gowebcore’s Base so we inherit every common knob
// (HTTP/TLS ports, CORS, logging, metrics, etc.).
type Config struct {
	config.Base `mapstructure:",squash"`

	// --- hub‑specific settings (expand later) ---
	MongoURI string `mapstructure:"mongo_uri"`
}
