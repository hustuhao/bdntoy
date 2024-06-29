package requester

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/goccy/go-json"

	jsoniter "github.com/json-iterator/go"
)

// httpPost does a HTTP POST request with a body, checks the response to be a 200 OK and returns it
func (c *Client) httpPost(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New("bad status: " + resp.Status)
	}

	return resp, nil
}

// httpPostBodyBytes reads the whole HTTP body and returns it
func (c *Client) httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := c.httpPost(ctx, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// httpDo sends an HTTP request and returns an HTTP response.
func (c *Client) httpDo(req *http.Request) (*http.Response, error) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	res, err := client.Do(req)

	log := slog.With("method", req.Method, "url", req.URL)

	if err != nil {
		log.Debug("HTTP request failed", "error", err)
	} else {
		log.Debug("HTTP request succeeded", "status", res.Status)
	}

	return res, err
}

//HTTPGet 简单实现 http 访问 GET 请求
func (c *Client) HTTPGet(urlStr string) ([]byte, error) {
	res, err := c.HTTPClient.Get(urlStr)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

// Headers return the HTTP Headers of the url
func (c *Client) Headers(url string) (http.Header, error) {
	res, err := c.Req(http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	return res.Header, nil
}

// Size get size of the url
func (c *Client) Size(url string) (int, error) {
	h, err := c.Headers(url)
	if err != nil {
		return 0, err
	}
	s := h.Get("Content-Length")
	if s == "" {
		return 0, errors.New("Content-Length is not present")
	}
	size, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// Req 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// urlStr (网址),
// post (post 数据),
// header (header 请求头数据), 进行网站访问。
// 返回值分别为 *http.Response, 错误信息
func (c *Client) Req(method string, urlStr string, post interface{}, header map[string]string) (*http.Response, error) {
	var (
		req   *http.Request
		obody io.Reader
	)

	if post != nil {
		switch value := post.(type) {
		case io.Reader:
			obody = value
		case map[string]string, map[string]int, map[string]interface{}, []int, []string:
			postData, err := jsoniter.Marshal(value)
			if err != nil {
				return nil, err
			}
			header["Content-Type"] = "application/json"
			obody = bytes.NewReader(postData)
		case string:
			obody = strings.NewReader(value)
		case []byte:
			obody = bytes.NewReader(value)
		default:
			return nil, fmt.Errorf("request.Req: unknow post type: %s", post)
		}
	}

	req, err := http.NewRequest(method, urlStr, obody)
	if err != nil {
		return nil, err
	}

	// 设置浏览器标识
	req.Header.Set("User-Agent", c.UserAgent)

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	return c.HTTPClient.Do(req)
}

// Fetch 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// urlStr (网址),
// post (post 数据),
// header (header 请求头数据), 进行网站访问。
// 返回值分别为 []byte, 错误信息
func (c *Client) Fetch(method string, urlStr string, post interface{}, header map[string]string) ([]byte, error) {
	res, err := c.Req(method, urlStr, post, header)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}
