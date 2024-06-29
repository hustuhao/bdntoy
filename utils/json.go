package utils

import (
	"io"

	"github.com/goccy/go-json"
)

//UnmarshalReader 将 r 中的 json 格式的数据, 解析到 v
func UnmarshalReader(r io.Reader, v interface{}) error {
	d := json.NewDecoder(r)
	return d.Decode(v)
}

//UnmarshalJSON 将 r 中的 json 格式的数据, 解析到 v
func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
