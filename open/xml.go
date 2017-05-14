package open

import "encoding/xml"

type XmlBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:",omitempty"`
	FromUserName string   `xml:",omitempty"`
	CreateTime   int64
	MsgType      string       `xml:",omitempty"`
	Content      string       `xml:",omitempty"`
	Event        string       `xml:"Event,omitempty"`
	EventKey     string       `xml:"EventKey,omitempty"`
	Image        *XmlImage    `xml:",omitempty"`
	Video        *XmlVideo    `xml:",omitempty"`
	Music        *XmlMusic    `xml:",omitempty"`
	ArticleCount int          `xml:",omitempty"`
	Articles     *XmlArticles `xml:"Articles,omitempty"`
}
type XmlArticles struct {
	Items []*XmlArticle `xml:",omitempty" json:"articles"`
}
type XmlImage struct {
	MediaId string
}
type XmlVideo struct {
	MediaId     string
	Title       string
	Description string
}
type XmlMusic struct {
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
}

type XmlArticle struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PicUrl      string   `json:"picurl"`
	Url         string   `json:"url"`
}
