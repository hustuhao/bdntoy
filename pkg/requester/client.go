package requester

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Client offers methods to download video metadata and video streams.
type Client struct {
	// HTTPClient can be used to set a custom HTTP client.
	// If not set, http.DefaultClient will be used
	//HTTPClient *http.Client
	HTTPClient *http.Client

	UserAgent string
}

//NewHTTPClient new client
func NewHTTPClient() *Client {
	rc := retryablehttp.NewClient()
	httpClient := rc.StandardClient()
	rc.RetryMax = 3

	rc.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retryNumber int) {
		if retryNumber > 0 {
			logger.Printf("api: %s, retry count:%d", req.RequestURI, retryNumber)
		}
	}
	c := &Client{
		HTTPClient: httpClient,
	}
	c.SetTimeout(10 * time.Second)
	c.ResetCookieJar()
	return c
}

//SetCookies 设置 jar
func (c *Client) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.HTTPClient.Jar.SetCookies(u, cookies)
}

func (c *Client) Cookies(u *url.URL) map[string]string {
	cookies := c.HTTPClient.Jar.Cookies(u)

	cstr := map[string]string{}

	for _, cookie := range cookies {
		cstr[cookie.Name] = cookie.Value
	}

	return cstr
}

//ResetCookieJar 重置cookie jar
func (c *Client) ResetCookieJar() {
	c.HTTPClient.Jar, _ = cookiejar.New(nil)
}

//SetTimeout 设置超时时间
func (c *Client) SetTimeout(t time.Duration) {
	c.HTTPClient.Timeout = t
}
