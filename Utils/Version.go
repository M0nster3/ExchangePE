package Utils

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/gookit/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func Version(version string){
	date,_ := ioutil.ReadFile("Version.txt")
	dates := string(date)
	Bbody:=strings.Split(dates,"\n")
	for i :=0;i<len(Bbody);i++{
		if strings.Contains(Bbody[i],version){
			Bbody1:=strings.Split(Bbody[i],"        ")
			fmt.Println(color.LightCyan.Sprintf("        对应版本名称: %s \n",Bbody1[0]))
		}
	}
}
func Inter(){
	os.Remove("Version.txt")

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body,_ := ioutil.ReadAll(resp.Body)
	Bbody:=strings.Split(string(body),"<tr>")
	for i :=0;i<len(Bbody);i++{
		pattern := regexp.MustCompile("nal\">(.*?)</a><")
		pattern1 := regexp.MustCompile("ter;\">(.*?)</td>")
		pattern2 := regexp.MustCompile("<td>([0-9]\\.[0-9]\\.[0-9]+)</td>")
		pattern3 := regexp.MustCompile("<td>Exchange(.*?)</td>")
		a := pattern.FindString(Bbody[i])
		a1 := pattern1.FindString(Bbody[i])
		a2 := pattern2.FindString(Bbody[i])
		a3 := pattern3.FindString(Bbody[i])
		if a!=""{
			//a2 := a[5:len(a)-5]+"        "+a1[5:]
			b2 := a[5:len(a)-5]+"        "+a1[6:len(a1)-5]+"\n"
			//println(a1)
			//data := []byte(a[5:len(a)-5])
			data := []byte(b2)
			// 检测文件是否存在
			if _, err := os.Stat("Version.txt"); err == nil {
				// 文件存在，打开文件并追加数据
				f, err := os.OpenFile("Version.txt", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return
				}
				defer f.Close()
				if _, err := f.Write(data); err != nil {
					return
				}
			} else {
				// 文件不存在，创建文件并写入数据
				fileee,_:=os.Create("Version.txt")
				defer func(fileee *os.File) {
					err := fileee.Close()
					if err != nil {
						return
					}
				}(fileee)
				buff:=bufio.NewWriter(fileee)
				buff.WriteString(a[5:len(a)-5]+"        "+a1[6:len(a1)-5]+"\n")
			}

		}else if a3!=""{
			if a1 !=""{
				b2 := a3[4:len(a3)-5]+"        "+a1[6:len(a1)-5]+"\n"
				data := []byte(b2)
				// 检测文件是否存在
				if _, err := os.Stat("Version.txt"); err == nil {
					// 文件存在，打开文件并追加数据
					f, err := os.OpenFile("Version.txt", os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						return
					}
					defer f.Close()
					if _, err := f.Write(data); err != nil {
						return
					}
				} else {
					// 文件不存在，创建文件并写入数据
					fileee,_:=os.Create("Version.txt")
					defer fileee.Close()
					buff:=bufio.NewWriter(fileee)
					buff.WriteString(a3[4:len(a3)-5]+"        "+a1[6:len(a1)-5]+"\n")
				}
			}else if a2 !="" {
				b2 := a3[4:len(a3)-5]+"        "+a2[4:len(a2)-5]+"\n"
				data := []byte(b2)
				// 检测文件是否存在
				if _, err := os.Stat("Version.txt"); err == nil {
					// 文件存在，打开文件并追加数据
					f, err := os.OpenFile("Version.txt", os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						return
					}
					defer f.Close()
					if _, err := f.Write(data); err != nil {
						return
					}
				} else {
					// 文件不存在，创建文件并写入数据
					fileee,_:=os.Create("Version.txt")
					defer fileee.Close()
					buff:=bufio.NewWriter(fileee)
					buff.WriteString(a3[4:len(a3)-5]+"        "+a2[4:len(a2)-5]+"\n")
				}
			}
		}
	}
}