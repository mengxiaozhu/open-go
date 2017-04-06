package open

type SignBody struct {
	MediaID   string `json:"media_id" sign:"true"`
	ApiKey    string `json:"api_key" sign:"true"`
	Timestamp string `json:"timestamp" sign:"true"`
	NonceStr  string `json:"nonce_str" sign:"true"`
	Sign      string `json:"sign"`
}

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type HookOpenReq SignBody
type HookCloseReq SignBody
type HookKeywordsReq struct {
	SignBody
	Keyword string `json:"keyword"`
}

type HookOpenResp struct {
	BaseResp
	Url      string `json:"url"`
	Token    string `json:"token"`
	IsConfig int    `json:"is_config"`
}

type HookKeywordResp BaseResp
type HookCloseResp BaseResp

type GetMediaInfoReq SignBody
type GetMediaKeywordsReq SignBody

type MediaInfo struct {
	Name        string `json:"name"`
	MediaNumber string `json:"media_number"`
	AvatarImage string `json:"avatar_image"`
	MediaID     string `json:"media_id"`
	MediaType   string `json:"media_type"`
	MediaURL    string `json:"media_url"`
	SchoolName  string `json:"school_name"`
	SchoolCode  string `json:"school_code"`
	VerifyType  string `json:"verify_type"`
}

type MediaKeywords struct {
	Name        string `json:"name"`
	MediaNumber string `json:"media_number"`
	AvatarImage string `json:"avatar_image"`
	MediaID     string `json:"media_id"`
	MediaURL    string `json:"media_url"`
	Keyword     string `json:"keyword"`
}
