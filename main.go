package main

import (
	"math/rand"
	"os"
	"time"

	cmd "github.com/eminmuhammadi/memcache/cmd"
	"github.com/urfave/cli"
)

var VERSION = "1.0.1-dev"
var BUILD_ID = "0"
var BUILD_TIME = "0"

var Commands = []cli.Command{
	cmd.Start(),
}

var Authors = []cli.Author{
	{
		Name:  "Emin Muhammadi",
		Email: "muemin.info@gmail.com",
	},
}

const banner = `
                                               _
                                              | |
 _ __ ___    ___  _ __ ___    ___  __ _   ___ | |__    ___ 
| '_   _ \  / _ \| '_   _ \  / __|/ _  | / __|| '_ \  / _ \
| | | | | ||  __/| | | | | || (__| (_| || (__ | | | ||  __/
|_| |_| |_| \___||_| |_| |_| \___|\__,_| \___||_| |_| \___|
`

func main() {
	rand.Seed(time.Now().UnixNano())

	if os.Getenv("MEMCACHE_DISABLE_BANNER") == "" {
		println(banner)
	}

	app := &cli.App{
		Name:      "memcache",
		Usage:     "http/https based in memory cache server written in Golang.",
		Version:   VERSION,
		Copyright: "memcache  Copyright (C) 2022  Emin Muhammadi",
		Authors:   Authors,
		ExtraInfo: func() map[string]string {
			return map[string]string{
				"LICENSE":    "The GNU General Public License",
				"VERSION":    VERSION,
				"BUILD":      BUILD_ID,
				"BUILD_TIME": BUILD_TIME,
			}
		},
		Commands:             Commands,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
