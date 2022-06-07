package core

import (
	"net"
)

func MX(domain string) (string, error) {
	res, err := net.LookupMX(domain)
	if len(res) != 0 {
		return res[0].Host, err
	}
	return "", err
}

//递归处理CNAME链
func CNAME(domain string) (string, error) {
	res, err := net.LookupCNAME(domain)
	if err == nil {
		res1, err1 := net.LookupCNAME(res)
		if err1 == nil && res != res1 { //不知道为什么github.io的CNAME值指向自己
			return CNAME(res1)
		} else {
			return res, err
		}
	} else {
		return res, err
	}
}

func NS(domain string) ([]*net.NS, error) {
	return net.LookupNS(domain)
}

func resolve(domain string) (t, dst string, err error) {
	//print(domain)
	dst, err = MX(domain)
	if dst != "" {
		return "MX", dst, err
	}

	dst, err = CNAME(domain)
	if err == nil {
		return "CNAME", dst, err
	}

	return "None", "", nil
}
