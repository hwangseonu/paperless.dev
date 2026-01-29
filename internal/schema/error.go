package schema

type Error struct {
	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}
