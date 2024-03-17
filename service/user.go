package service

type BdnLoginStatusRsp struct {
	Errno     int `json:"errno"` // -6: 无效登录，请重新登录
	LoginInfo struct {
		Bdstoken    string `json:"bdstoken"`
		PhotoUrl    string `json:"photo_url"`
		Svip10Id    string `json:"svip10_id"`
		Uk          int    `json:"uk"`
		UkStr       string `json:"uk_str"`
		Username    string `json:"username"`
		VipIdentity string `json:"vip_identity"`
		VipLevel    int    `json:"vip_level"`
		VipPoint    int    `json:"vip_point"`
		VipType     string `json:"vip_type"`
	} `json:"login_info"`
	Newno     string `json:"newno"`
	RequestId int64  `json:"request_id"`
	ShowMsg   string `json:"show_msg"`
}

func (s *Service) BdnLoginStatus() (*BdnLoginStatusRsp, error) {
	body, err := s.requestLoginStatus()

	if err != nil {
		return nil, err
	}

	defer body.Close()

	rsp := new(BdnLoginStatusRsp)
	if err := handleJSONParse(body, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
