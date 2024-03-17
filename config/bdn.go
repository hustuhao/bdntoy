package config

import (
	"net/http"

	"turato.com/bdntoy/service"
)

// BdnInfo 百度网盘登录信息
type BdnInfo struct {
	User
	Bduss  string `json:"bduss"`
	Stoken string `json:"stoken"`

	Gids  []string `json:"gid_list"`
	FsIds []string `json:"fsid_list"`

	Gid string `json:"gid"`
}

type User struct {
	Bdstoken string `json:"bdstoken"`
	PhotoUrl string `json:"photo_url"`
	Uk       int    `json:"uk"`
	Username string `json:"username"`
}

func (info *BdnInfo) IsValid() bool {
	return info.Bduss != "" && info.Stoken != ""
}

//Service geek time service
func (info *BdnInfo) Service() *service.Service {
	ser := service.NewService(info.Bduss, info.Stoken)

	return ser
}

//SetBdnInfoByBdussAndStoken 设置用户
func (c *ConfigsData) SetBdnInfoByBdussAndStoken(bduss, stoken string) (*BdnInfo, error) {
	s := service.NewService(bduss, stoken)
	bdnInfo := &BdnInfo{
		Bduss:  bduss,
		Stoken: stoken,
	}
	rsp, err := s.BdnLoginStatus()
	if err != nil {
		return nil, err
	}
	user := User{
		Bdstoken: rsp.LoginInfo.Bdstoken,
		PhotoUrl: rsp.LoginInfo.PhotoUrl,
		Uk:       rsp.LoginInfo.Uk,
		Username: rsp.LoginInfo.Username,
	}
	bdnInfo.User = user
	c.BdnInfo = *bdnInfo
	return bdnInfo, nil
}

//SetBdnInfoByCookies 设置用户
func (c *ConfigsData) SetBdnInfoByCookies(cookies string) (*BdnInfo, error) {
	// 解析cookie
	cs := cookieHeader(cookies)
	var BDUSS string
	var STOKEN string
	for _, v := range cs {
		if v.Name == "BDUSS" {
			BDUSS = v.Value
			continue
		}
		if v.Name == "STOKEN" {
			STOKEN = v.Value
			continue
		}
	}
	bdnInfo := &BdnInfo{
		Bduss:  BDUSS,
		Stoken: STOKEN,
	}
	s := bdnInfo.Service()
	rsp, err := s.BdnLoginStatus()
	if err != nil {
		return nil, err
	}
	user := User{
		Bdstoken: rsp.LoginInfo.Bdstoken,
		PhotoUrl: rsp.LoginInfo.PhotoUrl,
		Uk:       rsp.LoginInfo.Uk,
		Username: rsp.LoginInfo.Username,
	}
	bdnInfo.User = user
	c.BdnInfo = *bdnInfo
	return bdnInfo, nil
}

func cookieHeader(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}

//SetGidList 缓存gid列表
func (c *ConfigsData) SetGid(gid string) {
	c.BdnInfo.Gid = gid
	return
}

func (c *ConfigsData) GetGid() string {
	return c.BdnInfo.Gid
}

//SetGidList 缓存gid列表
func (c *ConfigsData) SetGidList(gidList []string) {
	c.BdnInfo.Gids = gidList
	return
}

//SetFsIdList 缓存fsid列表
func (c *ConfigsData) SetFsIdList(fsIdList []string) {
	c.BdnInfo.FsIds = fsIdList
	return
}

//ActiveService user service
func (c *ConfigsData) ActiveService() *service.Service {
	if c.service == nil {
		c.service = c.BdnInfo.Service()
	}
	return c.service
}
