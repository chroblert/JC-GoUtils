package jhttp

import (
	"strings"
)

type httpMsg struct {
	reqMethod   string
	reqHost     string
	reqUrl      string
	reqPath     string
	reqParams   map[string]string
	reqHeaders  map[string]string
	reqData     []byte
	isVerifySSL bool
	isUseSSL    bool

	intruData *intruderData
	//reqBytes []byte // 请求报文的字节数组
	//wordFiles []string // 字典文件切片
}
type intruderData struct {
	reqBytes  []byte   // 请求报文的字节数组
	wordFiles []string // 字典文件切片
}

func New() *httpMsg {
	return &httpMsg{
		reqMethod:   "Get",
		reqHost:     "",
		reqUrl:      "/",
		reqPath:     "/",
		reqParams:   make(map[string]string),
		reqHeaders:  make(map[string]string),
		reqData:     make([]byte, 0),
		isVerifySSL: false,

		intruData: &intruderData{
			reqBytes:  make([]byte, 0),
			wordFiles: make([]string, 0),
		},
	}
}

func (hm *httpMsg) InitWithFile(filename string) {
	hm.parseFromBurpReqFile(filename)
}

func (hm *httpMsg) InitWithBytes(reqMsg []byte) {
	hm.parseFromBytes(reqMsg)
}

func (hm *httpMsg) getInfoFromReqLine(reqLine []string) (reqMethod, reqPath string, reqParams map[string]string) {
	reqMethod = reqLine[0]
	reqParams = make(map[string]string)
	if strings.Index(reqLine[1], "?") != -1 {
		reqPath = reqLine[1][:strings.Index(reqLine[1], "?")]
		queryString := reqLine[1][strings.Index(reqLine[1], "?")+1:]
		for _, param := range strings.Split(queryString, "&") {
			idx := strings.Index(param, "=")
			reqParams[param[:idx]] = param[idx+1:]
		}
	} else {
		reqPath = reqLine[1]

	}
	return reqMethod, reqPath, reqParams
}

// 设置目标，如：http://test.test
func (hm *httpMsg) SetHost(target string) {
	if strings.Contains(target, "https") {
		hm.isVerifySSL = true
	}
	hm.reqHost = target[strings.Index(target, "/")+2:]

}

// 设置是否验证SSL
func (hm *httpMsg) SetIsVerifySSL(b bool) {
	hm.isVerifySSL = b
}

// 设置目标站点是否使用SSL
func (hm *httpMsg) SetIsUseSSL(b bool) {
	hm.isUseSSL = b
}

// 设置暴破用的字典所在的文件
func (hm *httpMsg) SetWordfiles(wordfiles ...string) {
	for _, v := range wordfiles {
		hm.intruData.wordFiles = append(hm.intruData.wordFiles, v)
	}
}

func (hm *httpMsg) Clean() {
	hm.reqMethod = "Get"
	hm.reqHost = ""
	hm.reqUrl = "/"
	hm.reqPath = "/"
	hm.reqParams = make(map[string]string)
	hm.reqHeaders = make(map[string]string)
	hm.reqData = make([]byte, 0)
	hm.isVerifySSL = false

	hm.intruData.reqBytes = make([]byte, 0)
	hm.intruData.wordFiles = make([]string, 0)
}
