package open

import (
	"fmt"
	"github.com/cocotyty/httpclient"
	"log"
	"time"
)

func (s *SDK) GetMediaInfoFromPlatform(openID string, platform string) (resp *MediaInfo, err error) {
	timestamp := fmt.Sprint(time.Now().Unix())
	req := &HookOpenReq{
		MediaID:   openID,
		ApiKey:    s.ApiKey,
		NonceStr:  s.nonceStr(),
		Timestamp: timestamp,
	}
	s.Sign(req)
	resp = &MediaInfo{}
	apiEntrypoint := "http://weixiao.qq.com/common"
	if platform == PlatformXiaozhu {
		apiEntrypoint = "http://www.mengxiaozhu.cn/open"
	}
	err = httpclient.Post(apiEntrypoint + "/get_media_info").
		Head("Content-Type", "application/x-www-form-urlencoded").
		JSON(req).
		Send().
		JSON(resp)
	log.Println(platform, req, err, resp)
	return

}
