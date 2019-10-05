package item

type VideoInformationTemp struct {
	Title          string   `json:"title" css_selector:"#title > b"`
	Code           string   `json:"code" css_selector:"#info > div:nth-of-type(2) > font > a"`
	PublishDate    string   `json:"publish_date" css_selector:"#info > div:nth-of-type(3)"`
	Length         string   `json:"length" css_selector:"#info > div:nth-of-type(4)"`
	Director       string   `json:"director" css_selector:"#info > div:nth-of-type(5) > a"`
	Maker          string   `json:"maker"  css_selector:"#info > div:nth-of-type(6) > a"`
	Issuer         string   `json:"issuer" css_selector:"#info > div:nth-of-type(7) > a"`
	Series         string   `json:"series" css_selector:"#info > div:nth-of-type(8) > a"`
	Category       []string `json:"category" css_selector:"#info > div:nth-of-type(9) > a"`
	Performer      string   `json:"performer" css_selector:"div.av_performer_name_box > a"`
	SmallImages    []string `json:"small_images" css_selector:"div.hvr-grow>a>img" attr:"src"`
	LargeImages    []string `json:"large_images" css_selector:"div.hvr-grow>a" attr:"href"`
	TorrentScripts []string `json:"torrent_scripts" css_selector:"body > div.dht_dl_area > script"`
	// Torrents       []Torrent `json:"torrents" css_selector:"body > div.dht_dl_area > div.dht_dl_title_content"`
	Torrents []string `json:"torrents"`
}

type Torrent struct {
	Name   string `json:"name" css_selector:"span > a"`
	Hd     bool   `json:"hd" css_selector:"span > a.HD_DL_icon_content"`
	Size   string `json:"size" css_selector:"div.dht_dl_size_content"`
	Magnet string `json:"magnet"`
}

type VideoInformation struct {
	Title       string    `json:"title"`
	Code        string    `json:"code"`
	PublishDate string    `json:"publish_date"`
	Length      string    `json:"length"`
	Director    string    `json:"director"`
	Maker       string    `json:"maker"`
	Issuer      string    `json:"issuer"`
	Series      string    `json:"series"`
	Category    []string  `json:"category"`
	Performer   string    `json:"performer"`
	SmallImages []string  `json:"small_images"`
	LargeImages []string  `json:"large_images"`
	Torrents    []Torrent `json:"torrents"`
}
