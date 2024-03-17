package service

type BdHistorySession struct {
	Errno     int                       `json:"errno"`
	RequestId int64                     `json:"request_id"`
	Records   []*BdHistorySessionRecord `json:"records"`
}

type BdHistorySessionRecord struct {
	Gid       string `json:"gid"` // 群id
	Name      string `json:"name"`
	Status    string `json:"status"`
	Time      int    `json:"time"`
	PhotoInfo []struct {
		Photo string `json:"photo"`
	} `json:"photoinfo"`
	MsgType   int      `json:"msg_type"`
	Gtype     string   `json:"gtype"`
	Uk        int64    `json:"uk"`
	FromUk    int64    `json:"from_uk"` // 分享者
	FromUname string   `json:"from_uname"`
	NickName  string   `json:"nick_name"`
	IsShare   int      `json:"is_share"`
	FileList  []string `json:"file_list"`
	Msg       string   `json:"msg"`
}

// GetHistorySession 获取历史会话记录
func (s *Service) GetHistorySession() ([]*BdHistorySessionRecord, error) {
	body, err := s.requestHistorySessions()
	if err != nil {
		return nil, err
	}
	defer body.Close()

	session := new(BdHistorySession)
	if err := handleJSONParse(body, session); err != nil {
		return nil, err
	}
	return session.Records, nil
}
