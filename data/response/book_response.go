package response

type BookResponse struct {
	Id       int    `json:"id"`
	BookName string `json:"bookname"`
	Author   string `json:"author"`
}
