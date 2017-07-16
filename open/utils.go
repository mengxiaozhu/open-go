package open

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type SDK struct {
	ApiKey    string
	ApiSecret string
}

func (s *SDK) nonceStr() string {
	return s.md5(strconv.FormatInt(time.Now().UnixNano(), 16))
}
func (s *SDK) md5(content string) (result string) {
	h := md5.New()
	h.Write([]byte(content))
	result = hex.EncodeToString(h.Sum(nil))
	return
}
func (s *SDK) calcSign(body interface{}) (signValue string) {
	params := url.Values{}
	v := reflect.ValueOf(body).Elem()
	k := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := k.Field(i)
		val := v.Field(i)
		tag := key.Tag
		jsonKey := tag.Get("json")
		_, sign := tag.Lookup("sign")
		if sign {
			switch val.Kind() {
			case reflect.String:
				params.Add(jsonKey, val.String())
			case reflect.Int:
				params.Add(jsonKey, strconv.Itoa(int(val.Int())))
			case reflect.Bool:
				params.Add(jsonKey, strconv.FormatBool(val.Bool()))
			}
		}
	}
	queryString := params.Encode()
	stringA := queryString
	stringSignTemp := stringA + "&" + "key=" + s.ApiSecret
	return strings.ToUpper(s.md5(stringSignTemp))
}

// 检查请求签名是否合法
func (s *SDK) CheckSign(body interface{}) (bool, error) {
	v := reflect.ValueOf(body).Elem()
	return v.FieldByName(SignField).String() == s.calcSign(body), nil
}

// 签名请求
func (s *SDK) Sign(body interface{}) {
	v := reflect.ValueOf(body).Elem()
	signValue := s.calcSign(body)
	log.Println(signValue)
	v.FieldByName(SignField).SetString(signValue)
}

type entrypoint struct {
	origin string
	u      *url.URL
}

func NewEntryPoint(raw string) (*entrypoint, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	return &entrypoint{
		origin: raw,
		u:      u,
	}, nil
}

func (ep *entrypoint) AddQuery(k string, v string) (*entrypoint) {
	query := ep.u.Query()
	query.Add(k, v)
	ep.u.RawQuery = query.Encode()
	return ep
}

func (ep *entrypoint) Build() (string) {
	return ep.u.String()
}
