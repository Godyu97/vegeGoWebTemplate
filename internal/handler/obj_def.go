package handler

import (
	"github.com/Godyu97/vegeGoWebTemplate/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所有200 的接口 resp 必须为ApiResp
type ApiResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty" swaggerignore:"true"`
}

// GetApi
type GetApi struct {
	UriToFnNameM map[string]string
}

func (a GetApi) SendOk(c *gin.Context, body any) {
	c.JSON(http.StatusOK, body)
}

func (a GetApi) SendBad(c *gin.Context, errMsg string) {
	//judge Code
	err := common.GetWebErrorFromErrMsg(errMsg)
	resp := ApiResp{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: nil,
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func (a GetApi) CheckAuth(c *gin.Context) error {
	return nil
}

func (a GetApi) UriToFnName(uri string) string {
	fnName, ok := a.UriToFnNameM[uri]
	if ok {
		return fnName
	} else {
		return uri
	}
}

// PostApi
type PostApi struct {
	UriToFnNameM map[string]string
}

func (a PostApi) SendOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func (a PostApi) SendBad(c *gin.Context, errMsg string) {
	//judge Code
	err := common.GetWebErrorFromErrMsg(errMsg)
	resp := ApiResp{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: nil,
	}
	c.JSON(http.StatusInternalServerError, resp)
}
func (a PostApi) CheckAuth(c *gin.Context) error {
	return nil
}

func (a PostApi) UriToFnName(uri string) string {
	fnName, ok := a.UriToFnNameM[uri]
	if ok {
		return fnName
	} else {
		return uri
	}
}
