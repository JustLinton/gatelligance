package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendPostRequest(addr string, body url.Values) string {
	resp, err := http.PostForm(addr, body)

	// resp, err := http.PostForm("http://114.215.149.199:8091/addLinkWork",
	// 	url.Values{"addr": {vl}})

	if err != nil {
		return "error"
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error"
	}

	return string(respBody)
}

// func CreateLinkTransaction(addr string) string {
// 	resp, err := http.PostForm("http://114.215.149.199:8091/addLinkWork",
// 		url.Values{"addr": {addr}})

// 	// print(addr)

// 	if err != nil {
// 		return "nil"
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "nil"
// 	}

// 	// fmt.Println(string(body))
// 	return string(body)
// }
