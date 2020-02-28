package details

type Book struct {
	Title            string   `json:"title"`
	ISBN             string   `json:"isbn"`
	PageCount        int      `json:"pageCount"`
	ThumbnailUrl     string   `json:"thumbnailUrl"`
	ShortDescription string   `json:"shortDescription"`
	LongDescription  string   `json:"longDescription"`
	Status           string   `json:"status"`
	Authors          []string `json:"authors"`
	Categories       []string `json:"categories"`
}
