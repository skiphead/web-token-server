package docs

type DocStruct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Scheme      string `json:"scheme"`
	Function    []struct {
		Description string `json:"description"`
		Method      string `json:"method"`
		Url         string `json:"url"`
		Security    bool   `json:"security"`
		Parameters  struct {
			Name      string `json:"name"`
			Body      string `json:"body"`
			Responses struct {
				ContentType string `json:"content_type"`
				Body        string `json:"body"`
				Code        struct {
					Ok       string `json:"ok"`
					NotFound string `json:"page_not_found"`
				} `json:"code"`
			} `json:"responses"`
		} `json:"parameters"`
	} `json:"function"`
	Version string `json:"version"`
	Date    string `json:"date"`
}
