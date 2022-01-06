package main

import (
	"github.com/chroblert/jgoutils/jlog"
	"github.com/chroblert/jgoutils/jrequests"
)

func main() {
	req, _ := jrequests.New()
	//req.SetIsVerifySSL(false)
	req.SetProxy("http://localhost:8080")
	req.SetIsVerifySSL(false)
	req.SetHttpVersion(2)
	req.SetKeepalive(false)
	_, err := req.Get("https://ipinfo.io")
	if err != nil {
		jlog.Error(err)
		return
	}
	req.SetHttpVersion(1)
	_, err = req.Get("https://ipinfo.io")
	if err != nil {
		jlog.Error(err)
		return
	}
	//req.SetHttpVersion(2)
	_, err = req.Get("https://myip.ipip.net")
	if err != nil {
		jlog.Error(err)
		return
	}
	//jlog.Info(resp.Resp.Header)
	//jlog.Info(resp.Resp.ProtoMajor)
	//jlog.Info(string(resp.Body()))
	//var wg = &sync.WaitGroup{}
	//for i := 0; i < 1; i++ {
	//	wg.Add(1)
	//	go func(t int) {
	//		defer wg.Done()
	//		resp, err := req.Get("http://myip.ipip.net") //jrequests.CGet("http://myip.ipip.net/").CSetParams(map[string][]string{strconv.Itoa(t): {strconv.Itoa(t)}}).CSetData(nil).CSetProxy("http://localhost:8080").CSetParams(map[string][]string{"q1": {"v1 '\"", "v2"}}).CSetHeaders(map[string][]string{"Content-Type": {"application/json"}, "Accept": {"application/json"}}).CDo()
	//		if err != nil {
	//			jlog.Error(err)
	//			return
	//		}
	//		jlog.Info(strings.TrimSpace(string(resp.Body())))
	//	}(i)
	//}
	//wg.Wait()
	//req.Get("http://myip.ipip.net") //jrequests.CGet("http://myip.ipip.net/").CSetParams(map[string][]string{strconv.Itoa(t): {strconv.Itoa(t)}}).CSetData(nil).CSetProxy("http://localhost:8080").CSetParams(map[string][]string{"q1": {"v1 '\"", "v2"}}).CSetHeaders(map[string][]string{"Content-Type": {"application/json"}, "Accept": {"application/json"}}).CDo()
	//req.Get("http://ipinfo.io")     //jrequests.CGet("http://myip.ipip.net/").CSetParams(map[string][]string{strconv.Itoa(t): {strconv.Itoa(t)}}).CSetData(nil).CSetProxy("http://localhost:8080").CSetParams(map[string][]string{"q1": {"v1 '\"", "v2"}}).CSetHeaders(map[string][]string{"Content-Type": {"application/json"}, "Accept": {"application/json"}}).CDo()
	//resp, err = jrequests.CGet("http://myip.ipip.net").CAddParams(map[string]string{"qqq1": "fdsa"}).CSetData(nil).CSetProxy("http://localhost:8080").CSetParams(map[string][]string{"q1": {"v1 '\"", "v2"}}).CSetHeaders(map[string][]string{"Content-Type": {"application/json"}, "Acceptxxx": {"application/json"}}).CAddHeaders(map[string]string{"kkkk": "kdfjadlksjf"}).CAddParams(map[string]string{"fadsfas": "{fasdfsad"}).CDo()
	//if err != nil {
	//	jlog.Error(err)
	//	return
	//}
	//jrequests.CGet("http://myip.ipip.net").CAddParams(map[string]string{"qqq1": "fdsa"}).CSetData(nil).CSetProxy("http://localhost:8080").CSetParams(map[string][]string{"q1": {"v1 '\"", "v2"}}).CSetHeaders(map[string][]string{"Content-Type": {"application/json"}, "Acceptxxx": {"application/json"}}).CAddHeaders(map[string]string{"kkkk": "kdfjadlksjf"}).CAddParams(map[string]string{"fadsfas": "{fasdfsad"}).CDo()
	_, err = jrequests.CGet("https://ipinfo.io").CSetIsVerifySSL(false).CSetHttpVersion(2).CSetProxy("http://localhost:8080").CSetTimeout(3).CDo()
	if err != nil {
		jlog.Error(err)
		//jlog.NFatal("not find arcsight ArcMC in",target)
		return
	}
	//_,err = jrequests.CGet("https://10.101.2.14/dashboard/locales/en/translation.json").CSetIsVerifySSL(false).CSetHttpVersion(2).CSetProxy("http://localhost:8080").CSetTimeout(3).CDo()
	//if err != nil{
	//	jlog.Error(err)
	//	//jlog.NFatal("not find arcsight ArcMC in",target)
	//	return
	//}
	//jlog.Info(string(resp.Body()))
}
