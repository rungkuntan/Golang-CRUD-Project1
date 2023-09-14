package request

type BookCreateRequest struct {
	BookName string `validate:"required min=1,max=100" json:"bookname"`
	Author   string `validate:"required min=1,max=100" json:"author"`
}