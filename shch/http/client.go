package main

import (
	"fmt"
	"github.com/nahid/gohttp"
	"time"
)

func main() {

	opt:=gohttp.SetTimeout(time.Duration(2)*time.Second)

	req := gohttp.NewRequest(opt)

	resp, err := req.
		FormData(map[string]string{"name": "Nahid"}).
		Post("https://httpbin.org/post")

	if err != nil {
		panic(err)
	}

	if resp.GetStatusCode() == 200 {
		str,_ := resp.GetBodyAsString()
		fmt.Println(str)
	}

}
