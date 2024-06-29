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

//NewBdCommand login command
func NewBdCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "hs",
			Usage:     "获取会话列表",
			UsageText: appName + " hs [number]",
			Action:    sessionAction,
			Before:    authorizationFunc,
			After:     configSaveFunc,
		},
		{
			Name:      "gs",
			Usage:     "获取分享群组",
			UsageText: appName + " gs [number]",
			Action:    groupAction,
			Before:    authorizationFunc,
			After:     configSaveFunc,
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
	// 如果输入了参数index
	if c.NArg() > 0 {
		index, err := strconv.Atoi(c.Args().Get(0))
		if err != nil {
			return errors.New("请输入正确的序号")
		}
		if index > len(sessions)-1 {

			return errors.New("请输入正确的序号")
		}
		session := sessions[index]
		config.Instance.SetGid(session.Gid)
		return fileLibraryAction(c)
	}

	renderSessions(sessions)
	return nil
}

func groupAction(c *cli.Context) error {
	groups, err := application.ShareGroups()
	if err != nil {
		return err
	}
	// 如果输入了参数index
	if c.NArg() > 0 {
		index, err := strconv.Atoi(c.Args().Get(0))
		if err != nil {
			return errors.New("请输入正确的序号")
		}
		if index > len(groups)-1 {
			return errors.New("请输入正确的序号")
		}
		group := groups[index]
		config.Instance.SetGid(group.Gid)
		return fileLibraryAction(c)
	}

	renderGroups(groups)
	return nil
}

func fileLibraryAction(c *cli.Context) error {
	gid := config.Instance.GetGid()
	if gid == "" {
		return errors.New("请先选择群组或会话")
	}

	rsp, err := application.FileLibraries(gid)
	if err != nil {
		return err
	}
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
			Num:    999999,
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

func renderGroups(groups []*service.ShareGroupInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "群名", "创建者", "创建时间"})
	table.SetAutoWrapText(false)

	for i, v := range groups {
		table.Append([]string{strconv.Itoa(i), v.Name, v.Uname, time.Unix(v.Ctime, 0).Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func renderFileLibrary(list []*service.ShareGroupFileMsg) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "文件库名称", "分享者", "创建时间"})
	table.SetAutoWrapText(false)

	for i, v := range list {
		msgCTime, _ := strconv.ParseInt(v.MsgCtime, 10, 64)
		var fileName string
		if len(v.FileList) > 0 {
			fileName = v.FileList[0].ServerFilename
		}
		table.Append([]string{strconv.Itoa(i), fileName, v.Uname, time.UnixMilli(msgCTime).Format("2006-01-02 15:04:05")})
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
