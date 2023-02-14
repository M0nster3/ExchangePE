package main

import (
	"ExchangePE/Utils"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)
var (
	domain string
	domains string
	results = make(chan int)

)

func Usage(){
	banner := `


▄███▄      ▄  ▄█▄     ▄  █ ▄███▄      ▄     ▄▀  ▄███▄   █ ▄▄  ▄███▄   
█▀   ▀ ▀▄   █ █▀ ▀▄  █   █ █▀   ▀      █  ▄▀    █▀   ▀  █   █ █▀   ▀  
██▄▄     █ ▀  █   ▀  ██▀▀█ ██▄▄    ██   █ █ ▀▄  ██▄▄    █▀▀▀  ██▄▄    
█▄   ▄▀ ▄ █   █▄  ▄▀ █   █ █▄   ▄▀ █ █  █ █   █ █▄   ▄▀ █     █▄   ▄▀ 
▀███▀  █   ▀▄ ▀███▀     █  ▀███▀   █  █ █  ███  ▀███▀    █    ▀███▀   
        ▀              ▀           █   ██                 ▀           
                                                                      
Usage:  
	ExchangePE.exe -domain DC.com
	ExchangePE.exe -domains target.txt

Options:
`
	print(banner)
	flag.PrintDefaults()

}
func init() {
	flag.StringVar(&domain,"domain", "", "添加DC域名")
	flag.StringVar(&domains,"domains", "", "检测Target文档")
	flag.Usage = Usage
}
var wg sync.WaitGroup
func HttpParse(dom string,line string){
	wg.Add(1)
	line=line+"."+dom
	if DNSA(line) == true {
		line = "https://"+line+"/owa/"
		if line != "" && strings.Contains(line,"."){
			//禁止检测证书
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			//client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
			//CheckRedirect不进行重定向
			client := &http.Client{Timeout: 5 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}, Transport: tr}
			resp, err := client.Get(line)
			if err == nil{
				defer resp.Body.Close()
				header := resp.Header
				for key, value := range header {
					if strings.Contains(key, "X-Owa-Version") && strings.Join(value, "") !=""{
						Bvalue := strings.Join(value, "")
						fmt.Println(color.LightGreen.Sprintf(" URL链接: %s \n",line))
						fmt.Println(color.LightGreen.Sprintf(" 内部版本号: %s \n",Bvalue))
						Utils.Version(Bvalue)
					}else if strings.Contains(key, "X-Owa-Version") && strings.Join(value, "") =="" {
						fmt.Println(color.LightGreen.Sprintf(" URL链接: %s \n",line))
						fmt.Println(color.LightRed.Sprintf(" 未返回版本信息 \n"))
					}else{
						continue
					}
				}
				<-results
			}else {
				fmt.Println(color.LightRed.Sprintf(" URL链接: %s \n",line))
				fmt.Println(color.LightRed.Sprintf(" Timeout \n"))
				<-results
			}

		} else {
			<-results
		}
	}else {
		<-results
	}
	wg.Done()
}

func DNSA (url string) bool{
	//addrs, err := net.LookupHost(url)
	_, err := net.LookupHost(url)
	if err == nil {
		return true
	}else {
		return false
	}

}
func main(){
	flag.Parse()
	data, _ := ioutil.ReadFile("domain.txt")
	content := string(data)
	lines := strings.Split(content, "\n")
	if domain != ""{
		fmt.Println(color.LightYellow.Sprintf(" Target Domain: %s \n",domain))
		for i:=0;i<=len(lines)-1;i++{
			go HttpParse(domain,lines[i])
		}
		for i:=0;i<=len(lines)-1;i++{
			results<-i
		}
		wg.Wait()
	}else if domains!=""{
		data1, _ := ioutil.ReadFile(domains)
		content1 := string(data1)
		lines1 := strings.Split(content1, "\n")

		for l:=0;l<=len(lines1)-1;l++{
			fmt.Println(color.LightYellow.Sprintf(" Target Domain: %s \n",lines1[l]))
			for i:=0;i<=len(lines)-1;i++{
				go HttpParse(lines1[l],lines[i])
			}
			for i:=0;i<=len(lines)-1;i++{
				results<-i
			}
			wg.Wait()
		}
	}else {
		Usage()
	}
}
