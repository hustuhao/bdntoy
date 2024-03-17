package service

import (
	"errors"
	"io"

	"turato.com/bdntoy/utils"
)

type resultData []byte

func (rd *resultData) UnmarshalJSON(data []byte) error {
	*rd = data

	return nil
}

func (rd *resultData) String() string {
	return string(*rd)
}

type resultError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (re *resultError) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "[]" {
		str = "{}"
	}

	type rError resultError

	e := new(rError)
	err := utils.UnmarshalJSON([]byte(str), &e)
	if err != nil {
		return err
	}

	*re = resultError(*e)

	return nil
}

// Result 从百度服务器解析的数据结构
type Result struct {
	Code    int        `json:"errno"`
	Data    resultData `json:"data"`
	ShowMsg string     `json:"show_msg"`
}

func (r *Result) isSuccess() bool {
	return r.Code == 0
}

func handleJSONParse(reader io.Reader, v interface{}) error {
	result := new(Result)
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = utils.UnmarshalJSON(data, result)
	if err != nil {
		return err
	}

	if !result.isSuccess() {
		//未登录或者登录凭证无效
		if result.Code == -6 {
			return ErrNotLogin
		}
		return errors.New(result.ShowMsg)
	}

	err = utils.UnmarshalJSON(data, v)
	if err != nil {
		return err
	}
	return nil
}
