package storage

import (
	"time"

	"github.com/influxdata/influxdb/v2/toml"
	"github.com/influxdata/influxdb/v2/tsdb"
)

// Default configuration values.
const (
	DefaultRetentionInterval       = time.Hour
	DefaultSeriesFileDirectoryName = "_series"
	DefaultIndexDirectoryName      = "index"
	DefaultWALDirectoryName        = "wal"
	DefaultEngineDirectoryName     = "data"
)

// Config holds the configuration for an Engine.
type Config struct {
	Data tsdb.Config

	// Frequency of retention in seconds.
	RetentionInterval toml.Duration `toml:"retention-interval"`
}

// NewConfig initialises a new config for an Engine.
func NewConfig() Config {
	return Config{
		Data:              tsdb.NewConfig(),
		RetentionInterval: toml.Duration(DefaultRetentionInterval),
	}
}
