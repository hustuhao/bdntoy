package cmds

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"turato.com/bdntoy/cli/version"
	"turato.com/bdntoy/config"
	"turato.com/bdntoy/pkg/logger"
)

var (
	_debug         bool
	_info          bool
	_stream        string
	appName        = filepath.Base(os.Args[0])
	configSaveFunc = func(c *cli.Context) error {
		err := config.Instance.Save()
		if err != nil {
			return errors.New("保存配置错误：" + err.Error())
		}
		return nil
	}
	authorizationFunc = func(c *cli.Context) error {
		if config.Instance.BdnInfo.Uk <= 0 {
			return config.ErrNotLogin
		}

		return nil
	}
)

//NewApp cli app
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "解析百度网盘会话分享文件库中的文件列表"
	app.Version = fmt.Sprintf("%s", version.Version)
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "Turn on debug logs",
			Destination: &_debug,
		},
	}

	app.Before = func(c *cli.Context) error {
		if _debug {
			l := logger.SetLogLevel(slog.LevelDebug.String())
			slog.SetDefault(l)
		}
		return nil
	}

	return app
}

//DefaultAction default action
func DefaultAction(c *cli.Context) error {
	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	sc := &NewBdCommand()[0]
	if sc != nil {
		return sc.Run(c)
	}

	return nil
}
