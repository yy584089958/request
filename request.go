package request

import (
	"net/http"
	"sync"
	"net/http/cookiejar"
)

type Browser struct {
	Client http.Client
	mu     sync.Mutex
}

//打开浏览器
func OpenBrowser() (Browser) {
	jar, _ := cookiejar.New(nil)
	return Browser{
		Client: http.Client{
			Jar: jar,
		},
	}
}

//Browser 模拟浏览器 发送 请求对象 记录cookie
//可以并发发送 记录cookie加锁
var DefaultReqHeader = map[string]string{
	"User-Agent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
}

func (n Browser) Send(req *http.Request) (res *http.Response, err error) {
	for k,v := range DefaultReqHeader{
		req.Header.Add(k,v)
	}
	res, err = n.Client.Do(req)
	if err!=nil{
		return nil,err
	}
	cookies := res.Cookies()
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Client.Jar.SetCookies(req.URL, cookies)
	return
}
