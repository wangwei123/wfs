/**
 * Copyright 2017 wfs Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package main

import (
	"flag"

	. "github.com/wangwei-go/wfs/conf"
	"github.com/wangwei-go/wfs/httpserver"
)

func main() {
	ParseFlag()
	flag.Parse()
	httpserver.Start()
}
