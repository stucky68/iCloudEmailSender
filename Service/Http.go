package Service

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func httpGet(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func httpPost(url, bodyData string, header map[string]string) (string, http.Header, []*http.Cookie, error) {
	client := &http.Client{}

	body := bytes.NewBuffer([]byte(bodyData))
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", nil, nil, err
	}

	if header != nil {
		// 处理Host
		//if host, ok := header["Host"]; ok {
		//	req.Host = host
		//}

		for key := range header {
			req.Header.Set(key, header[key])
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return "", nil, nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil, nil,err
	}


	return string(b), res.Header, res.Cookies(), nil
}
