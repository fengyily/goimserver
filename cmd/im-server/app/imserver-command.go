/*
 * @Author: F1
 * @Date: 2021-09-27 20:43:49
 * @LastEditTime: 2021-09-27 21:07:44
 * @LastEditors: F1
 * @Description:
 *  *
 *  *				Description
 *  *
 * @FilePath: /im-server/cmd/im-server/app/imserver-command.go
 *
 */
package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewIMServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "imserver",
		Long: `The im server is a instant messaging server.`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("hello! welcome to im server.")
		},
	}
	return cmd
}
