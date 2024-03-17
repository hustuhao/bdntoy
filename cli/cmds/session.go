package cmds

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"turato.com/bdntoy/cli/application"
	"turato.com/bdntoy/config"
	"turato.com/bdntoy/service"
)

//NewSessionCommand login command
func NewSessionCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "hs",
			Usage:     "获取历史会话列表",
			UsageText: appName + " session",
			Action:    sessionAction,
			Before:    authorizationFunc,
			//After:     configSaveFunc,
		},
		{
			Name:      "fl",
			Usage:     "获取分享文库列表",
			UsageText: appName + " fl",
			Action:    fileLibraryAction,
			Before:    authorizationFunc,
			After:     configSaveFunc,
		},
		{
			Name:      "tree",
			Usage:     "打印文件库目录树",
			UsageText: appName + " tree",
			Action:    filesAction,
			Before:    authorizationFunc,
		},
	}
}

func sessionAction(c *cli.Context) error {
	sessions, err := application.Sessions()

	if err != nil {
		return err
	}

	//gids := make([]string, 0, len(sessions))
	//for i := range sessions {
	//	gids = append(gids, sessions[i].Gid)
	//}
	//config.Instance.SetGidList(gids)
	renderSessions(sessions)
	return nil
}

func fileLibraryAction(c *cli.Context) error {
	args := c.Parent().Args()
	sessionIndex, err := strconv.Atoi(args.Get(1))
	if err != nil {
		cli.ShowCommandHelp(c, "fl")
		return errors.New("请输入会话ID")
	}

	records, err := application.Sessions()
	if err != nil {
		return err
	}
	if sessionIndex > len(records)-1 {
		return errors.New("请输入正确的序号")
	}

	rsp, err := application.FileLibraries(records[sessionIndex].Gid)
	if err != nil {
		return err
	}

	application.Sessions()
	config.Instance.SetGid(records[sessionIndex].Gid)
	renderFileLibrary(rsp.Records.MsgList)

	return nil
}

func filesAction(c *cli.Context) error {
	args := c.Parent().Args()
	fsIdIndex, err := strconv.Atoi(args.Get(1))
	if err != nil {
		cli.ShowCommandHelp(c, "fl")
		return errors.New("请输入文件库ID")
	}

	if err != nil {
		return err
	}

	rsp, err := application.FileLibraries(config.GetGid())
	r := rsp.Records.MsgList[fsIdIndex]
	var list []*service.ShareRecord
	for i := range r.FileList {
		param := service.ShareInfoParam{
			FromUk: r.Uk,
			ToUk:   r.Uk,
			MsgId:  r.MsgId,
			Num:    1000,
			Page:   0,
			FsId:   r.FileList[i].FsId,
			Gid:    config.GetGid(),
		}
		res, err := application.AllFiles(param)
		if err != nil {
			return err
		}

		list = append(list, res...)

	}

	renderFiles(list)
	return nil
}

func renderSessions(sessions []*service.BdHistorySessionRecord) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "群名", "邀请进群者", "更新时间(不太确定含义)"})
	table.SetAutoWrapText(false)

	for i, v := range sessions {
		table.Append([]string{strconv.Itoa(i), v.Name, v.FromUname, time.Unix(int64(v.Time), 0).Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func renderFileLibrary(list []*service.ShareGroupFileMsg) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "文件库名称", "内容", "创建时间"})
	table.SetAutoWrapText(false)

	for i, v := range list {
		msgCTime, _ := strconv.ParseInt(v.MsgCtime, 10, 64)
		table.Append([]string{strconv.Itoa(i), v.Uname, v.MsgContent, time.UnixMilli(msgCTime).Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func renderFiles(list []*service.ShareRecord) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "文件", "大小/MB", "创建时间"})
	table.SetAutoWrapText(false)

	for i, v := range list {
		c, _ := v.ServerCtime.Int64()
		size, _ := v.Size.Int64()
		mSIze := size >> 20
		table.Append([]string{strconv.Itoa(i), v.Path, v.ServerFilename, fmt.Sprint(mSIze), time.Unix(c, 0).Format("2006-01-02 15:04:05")})
	}
	table.Render()
}
