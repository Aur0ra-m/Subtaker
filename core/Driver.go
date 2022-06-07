package core

import (
	"fmt"
	"strings"
	"sync"
)

type Options struct {
	Domain     string
	DictPath   string
	FingerPath string
	Thread     int
	Timeout    int
}

func Drive(options *Options) {
	domain := options.Domain
	dictPath := options.DictPath
	thread := options.Thread
	fingerPath := options.FingerPath
	timeout := options.Timeout

	cond := sync.NewCond(new(sync.Mutex))

	//1.1 初始化-读取子域名爆破字典
	names := GetSubNames(dictPath)
	fmt.Printf("共有 [ %d ] 个子域名待测试\n", len(names))

	//1.2 初始化-读取指纹库
	fingers := readFingers(fingerPath)

	var wg sync.WaitGroup

	//2.处理子域名
	for _, name := range names {
		go Handle(&wg, domain, name, fingers, timeout, &thread, cond)
	}

	wg.Wait()

}

func Handle(wg *sync.WaitGroup, domain, name string, fingers []Finger, timeout int, thread *int, cond *sync.Cond) {
	//同步机制实现
	DecreaseSemaphore(thread, cond)
	wg.Add(1)

	//过滤CNAME、MX记录的subdomain，并返回记录值
	URI := name + "." + domain
	Type, dst, _ := resolve(URI)

	if dst == "" {
		wg.Done()

		IncreaseSemaphore(thread, cond)
		return
	}

	//校验是否存在注册利用
	IsVulnable := Verify1(URI, dst[:len(dst)-1])
	if IsVulnable || Type != "CNAME" {
		wg.Done()

		IncreaseSemaphore(thread, cond)
		return
	}

	//指纹校验
	Verify2(URI, dst[:len(dst)-1], timeout, fingers)

	IncreaseSemaphore(thread, cond)

	wg.Done()
}

// Verify1 调用域名查询接口 看是否可被注册
func Verify1(original, domain string) bool {
	pieces := strings.Split(domain, ".")
	OriginalDomain := pieces[len(pieces)-2] + "." + pieces[len(pieces)-1]

	//获取查询api
	api := QueryAPI()
	checkURI := api.APIURL + OriginalDomain

	//调用api查询注册信息
	_, res := Get(checkURI, 2)

	//响应失败则重新获取api并查询
	for {
		if strings.Contains(res, "status") || strings.Contains(res, "<returncode>200</returncode>") {
			break
		}
		api := QueryAPI()
		checkURI := api.APIURL + OriginalDomain
		_, res = Get(checkURI, 1)
	}

	/*依据接口返回信息，判断是否可被注册*/
	if strings.Contains(res, api.success) {
		fmt.Println("[+]可利用(域名可注册)" + original + "-->" + domain)
		return true
	}
	return false
}

// Verify2 指纹识别校验-》是否存在间接域名接管利用
func Verify2(original, domain string, timeout int, fingers []Finger) {
	URL := "https://" + domain
	_, res := Get(URL, timeout) //看网速分析，部分外网访问很慢，需要加大数值
	finger := Finger{}
	for _, finger = range fingers {
		for _, cname := range finger.Cname {
			if strings.Contains(domain, cname) {
				if strings.Contains(res, finger.Fingerprint) {
					fmt.Println("[+]可利用:" + original + "-->" + domain)
				}
				return
			}
		}
	}
	return
}
