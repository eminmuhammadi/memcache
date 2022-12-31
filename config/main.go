package config

import "fmt"

type Config struct {
	// DBPath is the path to sqlite.
	DBPath string
	// Timezone is the timezone used for the database.
	Timezone string
	// LogLevel is the log level used for the database.
	LogLevel string
	// EnablePrintRoutes is the flag to enable print routes.
	EnablePrintRoutes bool
	// ReduceMemoryUsage is the flag to reduce memory usage.
	ReduceMemoryUsage bool
	// DisableStartupMessage is the flag to disable startup message.
	DisableStartupMessage bool
	// BodyLimit is the maximum request body size.
	BodyLimit int
	// Concurrency is the maximum number of concurrent connections.
	Concurrency int
	// ReadBufferSize is the maximum buffer size for reading.
	ReadBufferSize int
	// WriteBufferSize is the maximum buffer size for writing.
	WriteBufferSize int
	// ReadTimeout is the maximum duration for reading the entire request, including the body.
	ReadTimeout int
	// WriteTimeout is the maximum duration before timing out writes of the response.
	WriteTimeout int
	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
	IdleTimeout int
	// Prefork is the flag to enable prefork
	Prefork bool
}

// Database user name
var __USERNAME__ = RandomString()

// Database password
var __PASSWORD__ = RandomString()

// Database salt for hashing
var __SALT__ = RandomString()

// DefaultConfig is the default configuration.
var DefaultConfig = Config{
	DBPath: fmt.Sprintf(
		"file::memory:?mode=memory&cache=shared&_auth&_auth_user=%s&_auth_pass=%s&_auth_salt=%s&_auth_crypt=sha512",
		__USERNAME__,
		__PASSWORD__,
		__SALT__),
	Timezone:              "UTC",
	LogLevel:              "SILENT",
	EnablePrintRoutes:     false,
	ReduceMemoryUsage:     true,
	DisableStartupMessage: true,
	Prefork:               false,
	BodyLimit:             4 * 1024 * 1024,
	Concurrency:           256 * 1024,
	ReadBufferSize:        4096,
	WriteBufferSize:       4096,
	ReadTimeout:           15,
	WriteTimeout:          15,
	IdleTimeout:           60,
}

// It sets the database path.
func (c *Config) SetDBPath(dbPath string) {
	c.DBPath = dbPath
}

// It sets the timezone.
func (c *Config) SetTimezone(timezone string) {
	c.Timezone = timezone
}

// It sets the log level.
func (c *Config) SetLogLevel(logLevel string) {
	c.LogLevel = logLevel
}

// It sets the flag to enable print routes.
func (c *Config) SetEnablePrintRoutes(enablePrintRoutes bool) {
	c.EnablePrintRoutes = enablePrintRoutes
}

// It sets the flag to reduce memory usage.
func (c *Config) SetReduceMemoryUsage(reduceMemoryUsage bool) {
	c.ReduceMemoryUsage = reduceMemoryUsage
}

// It sets the flag to disable startup message.
func (c *Config) SetDisableStartMessage(disableStartupMessage bool) {
	c.DisableStartupMessage = disableStartupMessage
}

// It sets the maximum request body size.
func (c *Config) SetBodyLimit(bodyLimit int) {
	c.BodyLimit = bodyLimit
}

// It sets the maximum buffer size for reading.
func (c *Config) SetReadBufferSize(readBufferSize int) {
	c.ReadBufferSize = readBufferSize
}

// It sets the maximum buffer size for writing.
func (c *Config) SetWriteBufferSize(writeBufferSize int) {
	c.WriteBufferSize = writeBufferSize
}

// It sets the maximum duration for reading the entire request, including the body.
func (c *Config) SetReadTimeout(readTimeout int) {
	c.ReadTimeout = readTimeout
}

// It sets the maximum duration before timing out writes of the response.
func (c *Config) SetWriteTimeout(writeTimeout int) {
	c.WriteTimeout = writeTimeout
}

// It sets the maximum amount of time to wait for the next request when keep-alives are enabled.
func (c *Config) SetIdleTimeout(idleTimeout int) {
	c.IdleTimeout = idleTimeout
}

// It sets the flag to enable prefork.
func (c *Config) SetPrefork(prefork bool) {
	c.Prefork = prefork
}

// It sets the maximum number of concurrent connections.
func (c *Config) SetConcurrency(concurrency int) {
	c.Concurrency = concurrency
}
