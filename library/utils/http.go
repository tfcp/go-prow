package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http请求
func RequestHttp(baseUrl string, query url.Values,
	postData, headers, cookies map[string]interface{},
	method string) (map[string]interface{}, error) {
	uri, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if query != nil {
		values := uri.Query()
		if values != nil {
			for k, v := range values {
				query[k] = v
			}
		}
		uri.RawQuery = query.Encode()
	}
	var postBody []byte
	if postData != nil {
		postBody, _ = json.Marshal(&postData)
	}
	req, err := http.NewRequest(strings.ToTitle(method), uri.String(), bytes.NewReader(postBody))
	if err != nil {
		return nil, err
	}
	// 添加header
	if headers != nil {
		for headerKey, headerValue := range headers {
			req.Header.Add(headerKey, headerValue.(string))
		}
	}
	// 添加cookie
	if cookies != nil {
		for cookieKey, cookieValue := range cookies {
			cookie := &http.Cookie{Name: cookieKey, Value: cookieValue.(string)}
			req.AddCookie(cookie)
		}
	}
	h := http.Client{}
	res, err := h.Do(req)
	if err != nil {
		return nil, err
	}
	rs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	rsMap := map[string]interface{}{}
	if err = json.Unmarshal(rs, &rsMap); err != nil {
		return nil, err
	}
	//if rsMap["retcode"].(float64) != 0 {
	//	return nil, errors.New(fmt.Sprintf("RequestHttp.error:%s", rs))
	//}
	return rsMap, err
}
