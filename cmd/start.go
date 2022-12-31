package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "github.com/eminmuhammadi/memcache/config"
	_dbPackage "github.com/eminmuhammadi/memcache/db"
	http "github.com/eminmuhammadi/memcache/http"
	"github.com/urfave/cli"
)

// Start starts the memcache server.
func Start() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "Starts the memcache server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "hostname",
				Usage:    "network interface to listen on",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Usage:    "network port to listen on",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "timezone",
				Value:    config.DefaultConfig.Timezone,
				Usage:    "timezone to use for time.Time",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "logLevel",
				Value:    config.DefaultConfig.LogLevel,
				Usage:    "log level to use; one of: info, warn, error, silent",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "bodyLimit",
				Value:    config.DefaultConfig.BodyLimit,
				Usage:    "in bytes (1024 * 1024 = 1MB)",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "readBufferSize",
				Value:    config.DefaultConfig.ReadBufferSize,
				Usage:    "in bytes",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "writeBufferSize",
				Value:    config.DefaultConfig.WriteBufferSize,
				Usage:    "in bytes",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "readTimeout",
				Value:    config.DefaultConfig.ReadTimeout,
				Usage:    "in seconds",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "writeTimeout",
				Value:    config.DefaultConfig.WriteTimeout,
				Usage:    "in seconds",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "idleTimeout",
				Value:    config.DefaultConfig.IdleTimeout,
				Usage:    "in seconds",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "prefork",
				Usage:    "enable preforking",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "concurrency",
				Value:    config.DefaultConfig.Concurrency,
				Usage:    "number of concurrent connections to handle",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "secure",
				Usage:    "enable TLS",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "certFile",
				Usage:    "path to TLS certificate file",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "keyFile",
				Usage:    "path to TLS key file",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			hostname := ctx.String("hostname")
			port := ctx.String("port")
			timezone := ctx.String("timezone")
			logLevel := ctx.String("logLevel")
			bodyLimit := ctx.Int("bodyLimit")
			readBufferSize := ctx.Int("readBufferSize")
			writeBufferSize := ctx.Int("writeBufferSize")
			readTimeout := ctx.Int("readTimeout")
			writeTimeout := ctx.Int("writeTimeout")
			idleTimeout := ctx.Int("idleTimeout")
			prefork := ctx.Bool("prefork")
			concurrency := ctx.Int("concurrency")
			secure := ctx.Bool("secure")
			certFile := ctx.String("certFile")
			keyFile := ctx.String("keyFile")

			if timezone != config.DefaultConfig.Timezone {
				config.DefaultConfig.SetTimezone(timezone)
			}

			if logLevel != config.DefaultConfig.LogLevel {
				config.DefaultConfig.SetLogLevel(logLevel)
			}

			if bodyLimit != config.DefaultConfig.BodyLimit {
				config.DefaultConfig.SetBodyLimit(bodyLimit)
			}

			if readBufferSize != config.DefaultConfig.ReadBufferSize {
				config.DefaultConfig.SetReadBufferSize(readBufferSize)
			}

			if writeBufferSize != config.DefaultConfig.WriteBufferSize {
				config.DefaultConfig.SetWriteBufferSize(writeBufferSize)
			}

			if readTimeout != config.DefaultConfig.ReadTimeout {
				config.DefaultConfig.SetReadTimeout(readTimeout)
			}

			if writeTimeout != config.DefaultConfig.WriteTimeout {
				config.DefaultConfig.SetWriteTimeout(writeTimeout)
			}

			if idleTimeout != config.DefaultConfig.IdleTimeout {
				config.DefaultConfig.SetIdleTimeout(idleTimeout)
			}

			if prefork {
				config.DefaultConfig.SetPrefork(prefork)
			}

			if concurrency != config.DefaultConfig.Concurrency {
				config.DefaultConfig.SetConcurrency(concurrency)
			}

			return startServer(secure, hostname, port, certFile, keyFile)
		},
	}
}

// Starts the server
func startServer(secure bool, hostname string, port string, certFile string, keyFile string) error {
	db, err := _dbPackage.Open()
	if err != nil {
		return err
	}

	db.AutoMigrate(&_dbPackage.Cache{})

	app := http.Create()

	go func() {
		println(fmt.Sprintf("%s Memcache is running on %s:%s", _dbPackage.TimeNowString(), hostname, port))

		if secure && certFile != "" && keyFile != "" {
			if err := http.StartSecure(app, db, hostname, port, certFile, keyFile); err != nil {
				panic(err)
			}
		} else {
			if err := http.Start(app, db, hostname, port); err != nil {
				panic(err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	ctx := make(chan os.Signal, 10)
	signal.Notify(ctx, os.Interrupt, syscall.SIGTERM)

	_signal := <-ctx
	println(fmt.Sprintf("%s Received %q signal", _dbPackage.TimeNowString(), _signal))

	if _signal == os.Interrupt || _signal == syscall.SIGTERM {
		app.Shutdown()

		println(fmt.Sprintf("%s Memcache is shutting down", _dbPackage.TimeNowString()))

		_dbPackage.Close(db)
	}

	return nil
}
