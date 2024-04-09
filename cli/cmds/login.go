package cmds

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"turato.com/bdntoy/config"
)

//Login login data
type Login struct {
	bduss   string
	stoken  string
	cookies string
}

//LoginConfig config
var LoginConfig Login

//NewLoginCommand login command
func NewLoginCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "login",
			Usage:     "登录百度网盘",
			UsageText: appName + " login [OPTIONS]",
			Action:    loginAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "bduss",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.bduss,
				},
				cli.StringFlag{
					Name:        "stoken",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.stoken,
				},
				cli.StringFlag{
					Name:        "cookies",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.cookies,
				},
			},
			After: configSaveFunc,
		},
	}
}

func loginAction(c *cli.Context) error {
	//通过 cookie 登录
	var (
		bduss   = LoginConfig.bduss
		stoken  = LoginConfig.stoken
		cookies = LoginConfig.cookies
		bdnInfo *config.BdnInfo
		err     error
	)

	if bduss != "" && stoken != "" {
		bdnInfo, err = config.Instance.SetBdnInfoByBdussAndStoken(bduss, stoken)
		if err != nil {
			return err
		}
		fmt.Printf("百度网盘登录验证登录成功, 昵称:%s,bdstoken:%s", bdnInfo.Username, bdnInfo.Bdstoken)
		return nil
	}
	if cookies != "" {
		bdnInfo, err = config.Instance.SetBdnInfoByCookies(cookies)
		if err != nil {
			return err
		}
		fmt.Printf("百度网盘登录验证登录成功, 昵称:%s,bdstoken:%s", bdnInfo.Username, bdnInfo.Bdstoken)
		return nil
	}

	return errors.New("请输入登录凭证信息")
}

func usersAction(c *cli.Context) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Uk", "昵称"})
	bdnInfo := config.Instance.BdnInfo
	table.Append([]string{strconv.Itoa(0), strconv.Itoa(bdnInfo.Uk), bdnInfo.Username})
	table.Render()
	return nil
}
