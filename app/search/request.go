package search

type RequestBody struct {
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
	Keywords   []string `json:"keywords"`
	ViewsFrom  int      `json:"viewsFrom"`
	ViewsTo    int      `json:"viewsTo"`
}
