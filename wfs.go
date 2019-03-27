/**
 * Copyright 2017 wfs Author. All Rights Reserved.
 * email: donnie4w@gmail.com
 */
package main

import (
	"flag"

	. "github.com/wangwei123/wfs/conf"
	"github.com/wangwei123/wfs/httpserver"
)

func main() {
	ParseFlag()
	flag.Parse()
	httpserver.Start()
}
