package config

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

var DefaultConfig = Config{
	DBPath:                "file::memory:?cache=shared&_pragma=foreign_keys(1)",
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

func (c *Config) SetDBPath(dbPath string) {
	c.DBPath = dbPath
}

func (c *Config) SetTimezone(timezone string) {
	c.Timezone = timezone
}

func (c *Config) SetLogLevel(logLevel string) {
	c.LogLevel = logLevel
}

func (c *Config) SetEnablePrintRoutes(enablePrintRoutes bool) {
	c.EnablePrintRoutes = enablePrintRoutes
}

func (c *Config) SetReduceMemoryUsage(reduceMemoryUsage bool) {
	c.ReduceMemoryUsage = reduceMemoryUsage
}

func (c *Config) SetDisableStartMessage(disableStartupMessage bool) {
	c.DisableStartupMessage = disableStartupMessage
}

func (c *Config) SetBodyLimit(bodyLimit int) {
	c.BodyLimit = bodyLimit
}

func (c *Config) SetReadBufferSize(readBufferSize int) {
	c.ReadBufferSize = readBufferSize
}

func (c *Config) SetWriteBufferSize(writeBufferSize int) {
	c.WriteBufferSize = writeBufferSize
}

func (c *Config) SetReadTimeout(readTimeout int) {
	c.ReadTimeout = readTimeout
}

func (c *Config) SetWriteTimeout(writeTimeout int) {
	c.WriteTimeout = writeTimeout
}

func (c *Config) SetIdleTimeout(idleTimeout int) {
	c.IdleTimeout = idleTimeout
}

func (c *Config) SetPrefork(prefork bool) {
	c.Prefork = prefork
}

func (c *Config) SetConcurrency(concurrency int) {
	c.Concurrency = concurrency
}
