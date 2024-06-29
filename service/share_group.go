package service

import (
	"fmt"

	"github.com/goccy/go-json"
)

type GetShareGroupFileListRsp struct {
	Errno       int                   `json:"errno"`
	RequestId   int64                 `json:"request_id"`
	LastMsgTime int64                 `json:"last_msg_time"`
	HasMore     int                   `json:"has_more"` // 是否还有下一页
	Timestamp   int                   `json:"timestamp"`
	Records     *ShareGroupFileRecord `json:"records"`
}

type ListShareRsp struct {
	Errno     int            `json:"errno"`
	RequestId int64          `json:"request_id"`
	Records   []*ShareRecord `json:"records"`
	Timestamp int            `json:"timestamp"`
	HasMore   int            `json:"has_more"` // has_more, int, 1:表示还有数据, 0:表示没有数据
}

type ShareRecord struct {
	Category       json.Number `json:"category"`
	FsId           json.Number `json:"fs_id"`
	Isdir          json.Number `json:"isdir"`
	LocalCtime     json.Number `json:"local_ctime"`
	LocalMtime     json.Number `json:"local_mtime"`
	Md5            string      `json:"md5"`
	Path           string      `json:"path"`
	ServerCtime    json.Number `json:"server_ctime"`
	ServerFilename string      `json:"server_filename"`
	ServerMtime    json.Number `json:"server_mtime"`
	Size           json.Number `json:"size"`
}

type ShareGroupFileRecord struct {
	MsgCount int                  `json:"msg_count"`
	MsgList  []*ShareGroupFileMsg `json:"msg_list"`
}

type ShareGroupFileMsg struct {
	MsgId      string           `json:"msg_id"`
	GroupId    string           `json:"group_id"`
	MsgStatus  int              `json:"msg_status"`
	MsgType    int              `json:"msg_type"`
	MsgCtime   string           `json:"msg_ctime"`
	FileList   []ShareGroupFile `json:"file_list"`
	MsgContent string           `json:"msg_content"`
	Uk         int64            `json:"uk"`
	AvatarUrl  string           `json:"avatar_url"`
	Uname      string           `json:"uname"`
	NickName   string           `json:"nick_name"`
}

type ShareGroupFile struct {
	BlockList      []string `json:"block_list"`
	Category       string   `json:"category"` // 文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	ExtentInt3     string   `json:"extent_int3"`
	ExtentTinyint1 string   `json:"extent_tinyint1"`
	ExtentTinyint2 string   `json:"extent_tinyint2"`
	ExtentTinyint3 string   `json:"extent_tinyint3"`
	ExtentTinyint4 string   `json:"extent_tinyint4"`
	FileKey        string   `json:"file_key"`
	FsId           string   `json:"fs_id"` // 文件在云端的唯一标识ID
	Isdelete       string   `json:"isdelete"`
	Isdir          string   `json:"isdir"`       // 是否目录，0 文件、1 目录
	LocalCtime     string   `json:"local_ctime"` //文件在客户端创建时间
	LocalMtime     string   `json:"local_mtime"` // 文件在客户端修改时间
	Md5            string   `json:"md5"`         // 文件的md5值，只有是文件类型时，该KEY才存在
	OperId         string   `json:"oper_id"`
	OwnerId        string   `json:"owner_id"`
	OwnerType      string   `json:"owner_type"`
	Path           string   `json:"path"`
	Privacy        string   `json:"privacy"`
	RealCategory   string   `json:"real_category"`
	ServerAtime    string   `json:"server_atime"`
	ServerCtime    string   `json:"server_ctime"` //文件在服务器创建时间
	ServerFilename string   `json:"server_filename"`
	ServerMtime    string   `json:"server_mtime"` // 文件在服务器修改时间
	Size           string   `json:"size"`         // 文件大小，单位B
	Status         string   `json:"status"`
	TkbindId       string   `json:"tkbind_id"`
	Videotag       string   `json:"videotag"`
	Wpfile         string   `json:"wpfile"`
}

type ListShareGroupRsp struct {
	Errno     int               `json:"errno"`
	RequestId int64             `json:"request_id"`
	Count     int               `json:"count"`
	Timestamp int               `json:"timestamp"`
	Records   []*ShareGroupInfo `json:"records"`
}

type ShareGroupInfo struct {
	Gid           string      `json:"gid"`
	Gnum          string      `json:"gnum"`
	Name          string      `json:"name"`
	Gdesc         string      `json:"gdesc"`
	Announce      string      `json:"announce"`
	Type          string      `json:"type"`
	Status        string      `json:"status"`
	Ctime         int64       `json:"ctime"`
	NameFlag      string      `json:"name_flag"`
	Vip           string      `json:"vip"`
	SpamCount     interface{} `json:"spam_count"`
	InviteStatus  string      `json:"invite_status"`
	SearchStatus  int         `json:"search_status"`
	BanpostStatus string      `json:"banpost_status"`
	CommonSwitch  string      `json:"common_switch"`
	Photoinfo     []struct {
		Uk    int64  `json:"uk"`
		Uname string `json:"uname"`
		Photo string `json:"photo"`
	} `json:"photoinfo"`
	Gtype       int    `json:"gtype"`
	GroupStatus int    `json:"group_status"`
	Uname       string `json:"uname"`
	Uk          int64  `json:"uk"`
	AvatarUrl   string `json:"avatar_url"`
	GroupRemark string `json:"group_remark"`
}

func (s *Service) GetShareGroups() ([]*ShareGroupInfo, error) {
	body, err := s.requestShareGroups()

	if err != nil {
		return nil, err
	}

	defer body.Close()

	rsp := new(ListShareGroupRsp)
	if err := handleJSONParse(body, &rsp); err != nil {
		return nil, err
	}
	return rsp.Records, nil
}

// GetShareGroupFileList 分享群文件列表
func (s *Service) GetShareGroupFileList(gid string) (*GetShareGroupFileListRsp, error) {
	body, err := s.requestShareGroupFileList(gid)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	rsp := new(GetShareGroupFileListRsp)
	if err := handleJSONParse(body, &rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type ShareFileParam struct {
	Uk      int64  `json:"uk"`
	MsgId   string `json:"msg_id"`
	GroupId string `json:"group_id"`
	FsId    string `json:"fs_id"`
}

func NewShareFileParam(uk int64, msgId, groupId, fsId string) *ShareFileParam {
	return &ShareFileParam{
		Uk:      uk,
		MsgId:   msgId,
		GroupId: groupId,
		FsId:    fsId,
	}
}

// GetShareFiles 分享群文件列表
func (s *Service) GetShareFiles(param ShareInfoParam) (*ListShareRsp, error) {
	body, err := s.requestShareInfo(param)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	rsp := new(ListShareRsp)
	if err := handleJSONParse(body, &rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (s *Service) GetAllShareFiles(param ShareInfoParam) ([]*ShareRecord, error) {
	param.Page = 0
	list, err := s.getAllShareFilesCore(param)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Service) getAllShareFilesCore(param ShareInfoParam) ([]*ShareRecord, error) {
	listShareRsp, err := s.GetShareFiles(param)
	if err != nil {
		return nil, err
	}
	var list []*ShareRecord
	for _, record := range listShareRsp.Records {
		if record.Isdir.String() == "0" { // 文件
			fmt.Printf("%s\n", record.Path)
			list = append(list, record)
		} else { // 目录
			param.FsId = record.FsId.String()
			records, err := s.getAllShareFilesCore(param) // 递归获取目录下的文件
			if err != nil {
				return nil, err
			}
			list = append(list, records...)
		}
	}

	return list, nil
}
