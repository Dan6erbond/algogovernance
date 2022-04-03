package client

// ActiveBanners represents any active banners to be shown by the official Algorand Governance website.
type ActiveBanners struct {
	ID              int    `json:"id"`
	DescriptionHTML string `json:"description_html"`
}
