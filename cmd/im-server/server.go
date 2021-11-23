/*
 * @Author: F1
 * @Date: 2021-09-27 20:37:55
 * @LastEditTime: 2021-09-27 21:09:41
 * @LastEditors: F1
 * @Description:
 *  *
 *  *				Description
 *  *
 * @FilePath: /im-server/cmd/im-server/server.go
 *
 */
package main

import (
	"im/cmd/im-server/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewIMServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
