package main

import (
	"encoding/json"
	"github.com/mengxiaozhu/open-go/open"
	"github.com/qiniu/log"
	"gopkg.in/macaron.v1"
	"time"
)

const (
	ApiKey    = "shared"
	ApiSecret = "shared"
)

var SDK *open.SDK = &open.SDK{ApiKey: ApiKey, ApiSecret: ApiSecret}

func openCtrl(ctx *macaron.Context, body []byte) {
	req := &open.HookOpenReq{}
	err := json.Unmarshal(body, req)
	if err != nil {
		ctx.JSON(200, &open.BaseResp{ErrCode: -1, ErrMsg: err.Error()})
		return
	}
	signed, err := SDK.CheckSign(req)
	if err != nil {
		ctx.JSON(200, &open.BaseResp{ErrCode: -1, ErrMsg: err.Error()})
		return
	}

	if !signed {
		ctx.JSON(200, &open.BaseResp{ErrCode: -1, ErrMsg: "invalid sign"})
		return
	}

	ctx.JSON(200, &open.HookOpenResp{
		Token: req.MediaID,
	})

	go func() {
		time.Sleep(5 * time.Second)
		SDK.GetMediaInfoFromPlatform(req.MediaID, open.PlatformXiaozhu)
	}()

}

func triggerCtrl(ctx *macaron.Context, body []byte) {
	target := ctx.Query("target")
	if target == "" {
		target = "https://www.mengxiaozhu.cn"
	}
	ctx.Redirect(target)
}
func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Any("/index", func(ctx *macaron.Context) {
		bodyBytes, err := ctx.Req.Body().Bytes()
		if err != nil {
			ctx.JSON(200, &open.BaseResp{ErrCode: -1, ErrMsg: err.Error()})
			return
		}
		log.Println("->", string(bodyBytes))

		typ := ctx.Query("type")
		switch typ {
		case "open":
			openCtrl(ctx, bodyBytes)
		case "trigger":
			triggerCtrl(ctx, bodyBytes)
		case "close":
			fallthrough
		case "keyword":
			ctx.JSON(200, &open.BaseResp{ErrCode: 0, ErrMsg: "ok"})
		}
		return
	})

	m.Run(4000)
}
