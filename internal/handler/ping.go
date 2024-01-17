package handler

import (
	"github.com/Godyu97/vegeGoWebTemplate/internal/common"
	"github.com/Godyu97/vegeGoWebTemplate/internal/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Godyu97/vegeGoWebTemplate/internal/types"
)

//	@Summary		Ping
//
//	@Description	Ping
//	@Tags			demo
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	ApiResp{data=types.PingResp}	"success"
//	@Failure		401	{object}	ApiResp{data=string}			"权限不足"
//	@Failure		500	{object}	ApiResp{}						"服务器内部错误"
//	@Router			/Ping [GET]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ApiResp{
		Code: common.GetWebMsgCode(common.Nil),
		Msg:  common.Nil,
		Data: types.PingResp{
			Field: "Pong pong pong",
		},
	})
}

//	@Summary		ObjPing
//
//	@Description	ObjPing
//	@Tags			demo
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	ApiResp{data=types.PingResp}	"success"
//	@Failure		401	{object}	ApiResp{data=string}			"权限不足"
//	@Failure		500	{object}	ApiResp{}						"服务器内部错误"
//	@Router			/v1/api/obj/ObjPing [GET]
func (a GetApi) ObjPing(ctx *gin.Context, req *struct{}) (ApiResp, error) {
	//logic 业务逻辑
	logic.DemoLogic()
	return ApiResp{
		Code: common.GetWebMsgCode(common.Nil),
		Msg:  common.Nil,
		Data: types.PingResp{
			Field: "Obj Pong pong pong",
		},
	}, nil
}
