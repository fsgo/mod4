/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/7/3
 */

package main

import (
	"fmt"

	"github.com/fsgo/mod1"
	"github.com/fsgo/mod2"
	"github.com/fsgo/mod3"
)

func main() {
	fmt.Println("mod2->mod1(v1.0.0)")
	fmt.Println("mod3->mod1(v1.0.1)")
	fmt.Println("mod4->mod1(v1.0.2)")

	fmt.Println("")

	fmt.Println("mod1.Version():", mod1.Version())

	fmt.Println("")

	fmt.Println("mod2.Version():", mod2.Version())

	fmt.Println("")

	fmt.Println("mod3.Version():", mod3.Version())
}
