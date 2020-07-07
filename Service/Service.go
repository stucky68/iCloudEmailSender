package Service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var g_XAppleIFDClientInfo XAppleIFDClientInfo
var g_Settings Settings

var g_UserAgent string


func GetAuthToken(account, password string) (token AuthToken, error error) {
	loginData := struct {
		AccountName string `json:"accountName"`
		Password string `json:"password"`
		RememberMe bool `json:"rememberMe"`
		TrustTokens []string `json:"trustTokens"`
	}{
		account,
		password,
		false,
		nil,
	}

	loginResult, err := json.Marshal(&loginData)
	if err != nil {
		return token, err
	}

	infoResult, err := json.Marshal(&g_XAppleIFDClientInfo)
	if err != nil {
		return token, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Referer"] = "https://idmsa.apple.com/appleauth/auth/signin"
	headers["Accept"] = "application/json, text/javascript, */*; q=0.01"
	headers["User-Agent"] = g_UserAgent
	headers["Origin"] = "https://idmsa.apple.com"
	headers["X-Apple-Widget-Key"] = g_Settings.XAppleWidgetKey
	headers["X-Apple-OAuth-Client-Id"] = g_Settings.XAppleWidgetKey
	headers["X-Requested-With"] = "XMLHttpRequest"
	headers["X-Apple-I-FD-Client-Info"] = string(infoResult)
	result, resHeader, _, err:= httpPost("https://idmsa.apple.com/appleauth/auth/signin?isRememberMeEnabled=true", string(loginResult), headers)
	token.Token = resHeader.Get("x-apple-session-token")
	token.SessionID = resHeader.Get("x-apple-id-session-id")
	token.Scnt = resHeader.Get("scnt")
	token.data = []byte(result)
	return token, nil
}

func AccountLogin(token AuthToken) (dsInfo DsInfo, err error) {
	accountData := struct {
		AccountCountryCode string `json:"accountCountryCode"`
		DsWebAuthToken     string `json:"dsWebAuthToken"`
		ExtendedLogin      bool   `json:"extended_login"`
	}{
		"USA",
		token.Token,
		false,
	}

	accountResult, err := json.Marshal(&accountData)
	if err != nil {
		return
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Referer"] = "https://www.icloud.com/"
	headers["Accept"] = "application/json, text/javascript, */*; q=0.01"
	headers["User-Agent"] = g_UserAgent
	headers["Origin"] = "https://www.icloud.com"

	var cookie []*http.Cookie
	result, _, cookie, err := httpPost("https://setup.icloud.com/setup/ws/1/accountLogin?clientBuildNumber="+
		g_Settings.ClientBuildNumber+
		"&clientMasteringNumber="+
		g_Settings.ClientMasteringNumber+
		"&clientId="+g_Settings.ClientId, string(accountResult), headers)

	err = json.Unmarshal([]byte(result), &dsInfo)
	if err != nil {
		return
	}

	dsInfo.Cookies = ""
	for index, _ := range cookie {
		if cookie[index].Name == "X-APPLE-WEBAUTH-TOKEN" || cookie[index].Name == "X-APPLE-WEBAUTH-USER" {
			dsInfo.Cookies += cookie[index].Name + "=" + "\"" + cookie[index].Value + "\"; "
		}

	}

	dsInfo.Cookies += " wmsid=44"
	return
}

func SaveDraft(info DsInfo) (saveDraftResult SaveDraftResult, err error){
	cur := time.Now()
	timestamp := cur.UnixNano() / 1000000

	draftData := SaveDraftData{
		Jsonrpc: "2.0",
		Id:      strconv.FormatInt(timestamp, 10) + "/2",
		Method:  "saveDraft",
		Params: struct {
			Date               string `json:"date"`
			From               string `json:"from"`
			WebmailClientBuild string `json:"webmailClientBuild"`
			TextBody           string `json:"textBody"`
			Body               string `json:"body"`
			Attachments []int `json:"attachments"`
		}{
			cur.UTC().Format(time.RFC1123Z),
			info.DsInfo.FullName + "<" + info.DsInfo.PrimaryEmail + ">",
			g_Settings.ClientMasteringNumber,
			"\n",
			"",
			[]int{},
		},
		UserStats:   struct{}{},
		SystemStats: []int{0, 0, 0, 0},
	}
	draftDataResult, err := json.Marshal(&draftData)
	if err != nil {
		return
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	headers["Content-Length"] = strconv.Itoa(len(draftDataResult))
	headers["Accept"] = "*/*"
	headers["User-Agent"] = g_UserAgent
	headers["Origin"] = "https://www.icloud.com"
	headers["Cache-Control"] = "no-cache"
	headers["Connection"] = "keep-alive"
	headers["Host"] = "p34-mailws.icloud.com"
	headers["Pragma"] = "no-cache"
	headers["Cookie"] = info.Cookies

	result, _, _, err := httpPost("https://p34-mailws.icloud.com/wm/message?clientBuildNumber="+
		g_Settings.ClientBuildNumber+
		"&clientMasteringNumber="+
		g_Settings.ClientMasteringNumber+
		"&clientId="+g_Settings.ClientId+
		"&dsid=" +info.DsInfo.Dsid, string(draftDataResult), headers)

	err = json.Unmarshal([]byte(result), &saveDraftResult)
	return
}

func SendEmail(info DsInfo, saveDraftResult SaveDraftResult) {
	cur := time.Now()
	timestamp := cur.UnixNano() / 1000000

	draftData := SendEmailData{
		Jsonrpc: "2.0",
		Id:      strconv.FormatInt(timestamp, 10) + "/2",
		Method:  "send",
		Params: struct {
			Date string `json:"date"`
			From string `json:"from"`
			WebmailClientBuild string `json:"webmailClientBuild"`
			To string `json:"to"`
			Subject string `json:"subject"`
			TextBody string `json:"textBody"`
			Body string `json:"body"`
			Attachments []int `json:"attachments"`
			DraftGuid string `json:"draft_guid"`
		}{
			cur.UTC().Format(time.RFC1123Z),
			info.DsInfo.FullName + "<" + info.DsInfo.PrimaryEmail + ">",
			g_Settings.ClientMasteringNumber,
			"6992917@qq.com",
			"测试邮件",
			"\n",
			"我是测试内容",
			[]int{},
			saveDraftResult.Result.Guid,
		},
		UserStats:   struct{}{},
		SystemStats: []int{0, 0, 0, 0},
	}
	dataResult, err := json.Marshal(&draftData)
	if err != nil {
		return
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	headers["Content-Length"] = strconv.Itoa(len(dataResult))
	headers["Accept"] = "*/*"
	headers["User-Agent"] = g_UserAgent
	headers["Origin"] = "https://www.icloud.com"
	headers["Cache-Control"] = "no-cache"
	headers["Connection"] = "keep-alive"
	headers["Host"] = "p34-mailws.icloud.com"
	headers["Pragma"] = "no-cache"
	headers["Cookie"] = info.Cookies

	result, _, _, err := httpPost("https://p34-mailws.icloud.com/wm/message?clientBuildNumber="+
		g_Settings.ClientBuildNumber+
		"&clientMasteringNumber="+
		g_Settings.ClientMasteringNumber+
		"&clientId="+g_Settings.ClientId+
		"&dsid=" +info.DsInfo.Dsid, string(dataResult), headers)

	err = json.Unmarshal([]byte(result), &saveDraftResult)
	fmt.Println(result)
	return
}

func init()  {
	g_UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.1 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.1"
	g_XAppleIFDClientInfo = XAppleIFDClientInfo{
		U: g_UserAgent,
		L: "en_US",
		Z: "GMT+02:00",
		V: "1.1",
		F: "F0a44j1e3NlY5BNlY5BSmHACVZXnN9Qg3eZWv4JKYfSHolk2dUf.j7J1gBZEMgzH_y37lY2U.6elV2pNK1e3uJuMukAm4.f282pujsFjn45BNlY5CGWY5BOgkLT0XxU..7yv",
	}
	
	g_Settings = Settings{
		Language:              "en-us",
		Locale:                "en_US",
		XAppleWidgetKey:       "d39ba9916b7251055b22c7f910e2ea796ee65e98b2ddecea8f5dde8d9d1a815d",
		Timezone:              "US/Pacific",
		ClientBuildNumber:     "2010B30",
		ClientMasteringNumber: "2010Project39",
		ClientId: "38cbcf58-a87a-4360-832b-ed98dd5451e8",
	}
}
