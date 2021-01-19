/*
 * Copyright(C) 2021 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2021/1/19
 */

package main

import (
	"flag"

	"fsgo/mod4/internal"
)

var msg = flag.String("msg", "hello", "")

func main() {
	flag.Parse()
	internal.Say(*msg)
}
