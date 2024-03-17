package service

import (
	"os"
	"testing"

	"turato.com/bdntoy/pkg/requester"
)

var (
	// 加载登录信息
	bduss      = ""
	stoken     = ""
	s          = NewService(bduss, stoken)
	recods     []*BdHistorySessionRecord
	filerecord *ShareGroupFileRecord
)

func TestMain(m *testing.M) {

	exitValue := m.Run()
	os.Exit(exitValue)
}

func TestMainOrder(t *testing.T) {

	t.Run("TestService_GetHistorySession", TestService_GetHistorySession)

	t.Run("TestService_GetShareGroupFileList", TestService_GetShareGroupFileList)

	t.Run("TestService_GetAllShareFiles", TestService_GetAllShareFiles)
}

func TestService_GetHistorySession(t *testing.T) {
	type fields struct {
		client *requester.Client
	}
	recods, _ = s.GetHistorySession()
	t.Log(recods)
}

func TestService_GetShareGroupFileList(t *testing.T) {
	for i := range recods {
		rsp, err := s.GetShareGroupFileList(recods[i].Gid)
		if i == 0 {
			filerecord = rsp.Records

		}
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rsp)
	}
}

func TestService_GetAllShareFiles(t *testing.T) {
	list := make([]string, 0)
	for _, v := range filerecord.MsgList {
		for _, f := range v.FileList {
			if f.Isdir == "0" {

			}
			pp := ShareInfoParam{
				FromUk: v.Uk,
				ToUk:   v.Uk,
				MsgId:  v.MsgId,
				Num:    100,
				Page:   0,
				FsId:   f.FsId,
				Gid:    v.GroupId,
			}
			rs, err := s.GetAllShareFiles(pp)
			if err != nil {
				t.Log(err)
			}
			for i := range rs {
				list = append(list, rs[i].Path)
			}
		}
	}
	t.Log(list)
}
