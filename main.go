package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"turato.com/bdntoy/cli/cmds"
	"turato.com/bdntoy/config"
	"turato.com/bdntoy/pkg/logger"

	"log/slog"
)

func init() {
	l := logger.SetLogLevel(os.Getenv("LOG_LEVEL"))
	slog.SetDefault(l)
	err := config.Instance.Init()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := cmds.NewApp()
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, cmds.NewLoginCommand()...)
	app.Commands = append(app.Commands, cmds.NewSessionCommand()...)

	app.Action = cmds.DefaultAction

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
