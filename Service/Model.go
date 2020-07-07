package Service

type XAppleIFDClientInfo struct {
	U string `json:"U"`
	L string `json:"L"`
	Z string `json:"Z"`
	V string `json:"V"`
	F string `json:"F"`
}

type Settings struct {
	Language string
	Locale string
	XAppleWidgetKey string
	Timezone string
	ClientBuildNumber string
	ClientMasteringNumber string
	ClientId string
}

type AuthToken struct {
	Token string
	SessionID string
	Scnt string
	data []byte
}

type DsInfo struct {
	DsInfo struct {
		Dsid string `json:"dsid"`
		PrimaryEmail string `json:"primaryEmail"`
		FullName string `json:"fullName"`
	} `json:"dsInfo"`

	Cookies string
}

type SaveDraftResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Id string `json:"id"`
	Result struct {
		Guid string `json:"guid"`
		Uid string `json:"uid"`
	}
}

type SendEmailData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id string `json:"id"`
	Method string `json:"method"`

	Params struct {
		Date string `json:"date"`
		From string `json:"from"`
		WebmailClientBuild string `json:"webmailClientBuild"`
		To string `json:"to"`
		Subject string `json:"subject"`
		TextBody string `json:"textBody"`
		Body string `json:"body"`
		Attachments []int `json:"attachments"`
		DraftGuid string `json:"draft_guid"`
	} `json:"params"`

	UserStats struct {
	}
	SystemStats []int
}

type SaveDraftData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id string `json:"id"`
	Method string `json:"method"`

	Params struct {
		Date string `json:"date"`
		From string `json:"from"`
		WebmailClientBuild string `json:"webmailClientBuild"`
		TextBody string `json:"textBody"`
		Body string `json:"body"`
		Attachments []int `json:"attachments"`
	} `json:"params"`

	UserStats struct {
	}
	SystemStats []int
}

type ListDraftData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id string `json:"id"`
	Method string `json:"method"`
	Params struct {
		Count int `json:"count"`
		Guid string `json:"guid"`
		Requesttype string `json:"requesttype"`
		Rollbackslot string `json:"rollbackslot"`
		Searchtext string `json:"body"`
		Searchtype interface{} `json:"searchtype"`
		Selected int `json:"selected"`
		Sortorder string `json:"sortorder"`
		Sorttype string `json:"sorttype"`
	} `json:"params"`
	UserStats struct {
	}
	SystemStats []int
}




