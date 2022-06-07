package core

type Finger struct {
	Service     string   `json:"service"`
	Cname       []string `json:"cname"`
	Fingerprint string   `json:"fingerprint"`
}
