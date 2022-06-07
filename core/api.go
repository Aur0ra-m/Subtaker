package core

import (
	"math/rand"
	"time"
)

type ApiInfo struct {
	APIURL  string
	success string
}

// QueryAPI 随机返回查询API,避免被封禁
func QueryAPI() ApiInfo {
	//随机种子-》随机数
	rand.Seed(time.Now().Unix())
	f := rand.Intn(2)

	if f == 0 {
		return ApiInfo{"https://www.cnz5.com/domain-registration/domain.php?action=caajax&domain_name=", "\"available\",\"price\""}
	}
	return ApiInfo{"http://panda.www.net.cn/cgi-bin/check.cgi?area_domain=", "<original>210"}
}
