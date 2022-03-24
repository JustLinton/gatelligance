package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetAudioSummary(addr string, id string) string {
	resp, err := http.PostForm("http://114.215.149.199:8091/addLinkWork",
		url.Values{"addr": {addr}, "id": {"123"}})

	print(addr + " " + id)

	if err != nil {
		return "nil"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "nil"
	}

	// fmt.Println(string(body))
	return string(body)
}
