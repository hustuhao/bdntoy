package service

import (
	"net/http"
	"net/url"

	"turato.com/bdntoy/pkg/requester"
)

var (
	geekBangCommURL = &url.URL{
		Scheme: "https",
		Host:   "pan.baidu.com",
	}
)

//Service bdnetdisk service
type Service struct {
	client *requester.Client
}

//NewService new service
func NewService(bduss, stoken string) *Service {
	client := requester.NewHTTPClient()
	client.ResetCookieJar()
	var cookies []*http.Cookie
	cookies = append(cookies, &http.Cookie{
		Name:   "BDUSS",
		Value:  bduss,
		Domain: "." + geekBangCommURL.Host,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "STOKEN",
		Value:  stoken,
		Domain: "." + geekBangCommURL.Host,
	})
	client.SetCookies(geekBangCommURL, cookies)
	return &Service{client: client}
}

//Cookies get cookies string
func (s *Service) Cookies() map[string]string {
	return s.client.Cookies(geekBangCommURL)
}
