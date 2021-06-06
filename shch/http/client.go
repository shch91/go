package http

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func main() {

	connectTimeout := 5 * time.Second
	readWriteTimeout := 100 * time.Millisecond

	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial:            TimeoutDialer(connectTimeout, readWriteTimeout),
		},
	}

	uri := "https://www.youtobe.com"
	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		fmt.Println("req error:" + err.Error())
		return
	}

	req.Header = http.Header{}

	req.Header["Connection"] = []string{"Close"}
	req.Header["User-Agent"] = []string{"wpt-http-client/1.1"}

	data := make(url.Values)

	data["name"] = []string{"baixs"}
	data["hobby"] = []string{"runing"}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("do error,err:" + err.Error())
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respBody))

}
