package Utils

import (
	"crypto/tls"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Version(version string){
	line:="https://learn.microsoft.com/zh-cn/exchange/new-features/build-numbers-and-release-dates?view=exchserver-2019"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	//CheckRedirect不进行重定向
	client := &http.Client{Timeout: 10 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}, Transport: tr}
	resp, _ := client.Get(line)
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	Bbody:=strings.Split(string(body),"<tr>")
	for i :=0;i<len(Bbody);i++{
		if strings.Contains(Bbody[i],version){
			pattern := regexp.MustCompile("nal\">(.*?)</a><")
			a := pattern.FindString(Bbody[i])
			fmt.Println(color.LightGreen.Sprintf(" 对应版本名称: %s \n\n",a[5:len(a)-5]))
		}
	}
	return
}