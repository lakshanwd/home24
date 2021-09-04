package model

type AnchorTag struct {
	Url        string `json:"url"`
	External   bool   `json:"external"`
	Accessible bool   `json:"accessible"`
}

type HtmlDocument struct {
	Host               string      `json:"-"`
	UserNameFieldExist bool        `json:"username"`
	PasswordFieldExist bool        `json:"password"`
	Title              string      `json:"title"`
	AnchorsTags        []AnchorTag `json:"aTags"`
	H1_Count           int         `json:"h1"`
	H2_Count           int         `json:"h2"`
	H3_Count           int         `json:"h3"`
	H4_Count           int         `json:"h4"`
	H5_Count           int         `json:"h5"`
	H6_Count           int         `json:"h6"`
}

type WebCrawlerRequest struct {
	Url string `json:"url"`
}
