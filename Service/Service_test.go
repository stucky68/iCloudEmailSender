package Service

import (
	"fmt"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	token, _ := GetAuthToken("iceynader0@icloud.com", "Az112211")
	info, _ := AccountLogin(token)
	result, _ := SaveDraft(info)
	fmt.Println(result)
	SendEmail(info, result)

	//url := "https://p34-mailws.icloud.com/wm/message?clientBuildNumber=2010B30&clientMasteringNumber=2010Project39&clientId=38cbcf58-a87a-4360-832b-ed98dd5451e8&dsid=17204903810"
	//method := "POST"
	//
	//payload := strings.NewReader("{\"jsonrpc\":\"2.0\",\"id\":\"1594132573059/2\",\"method\":\"list\",\"params\":{\"count\":50,\"guid\":\"folder:Sent Messages\",\"requesttype\":\"index\",\"rollbackslot\":\"0.0\",\"body\":\"\",\"searchtype\":\"\",\"selected\":1,\"sortorder\":\"descending\",\"sorttype\":\"Date\"},\"UserStats\":{},\"SystemStats\":[0,0,0,0]}")
	//
	//client := &http.Client {
	//}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//req.Header.Add("Cookie", "X-APPLE-WEBAUTH-TOKEN=\"v=2:t=HA==BST_IAAAAAAABLwIAAAAAF8EiZgRDmdzLmljbG91ZC5hdXRovQAnfrVcEJP11TfwCuj7q6XzonIQB1NlG5T5njuvlLJP36Bvoe0C6m83ERpvIEnX_qQIkCC49vHxC77P8e6y-5WDh-Z4NMHuj5jJy4ulue0OuDZ1rh9xWjymQkpe5lLP76NVuBPaL8nRPxYdOXsBOyf6wux7xw~~\"; X-APPLE-WEBAUTH-USER=\"v=1:s=0:d=17204903810\";; wmsid=44; X-APPLE-WEBAUTH-TOKEN=\"v=2:t=HA==BST_IAAAAAAABLwIAAAAAF8Eiv0RDmdzLmljbG91ZC5hdXRovQBV7Q3Xw4GnA-aGa1O-wgm8SdkYpa-PQ5Khy4iO0vPIetBjJ0pbNZZ-xwHwnGB-qITp2KQgZs2TriOl0bfJMZjFnKCtIfBe4FZ1R9i6ZeZjRJF-jrQwweL8_SEdOCKquZp6gIFSZ5VXmOzOkbwZucmwLp5rgw~~\"")
	//req.Header.Add("Origin", "https://www.icloud.com")
	//req.Header.Add("Content-Type", "text/plain")
	//
	//res, err := client.Do(req)
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(string(body))
}
