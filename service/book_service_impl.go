package service

import (
	"context"
	"go_project/data/request"
	"go_project/data/response"
	"go_project/errorhelper"
	"go_project/model"
	"go_project/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := model.Book{
		BookName: request.BookName,
		Author: request.Author,
	}
	b.BookRepository.Save(ctx, book)
}

// Delete implements BookService
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	errorhelper.PanicIfErr(err)
	b.BookRepository.Delete(ctx, book.Id)
}

// FindAll implements BookService
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, BookName: value.BookName,Author: value.Author}
		bookResp = append(bookResp, book)
	}
	return bookResp

}

// FindById implements BookService
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	errorhelper.PanicIfErr(err)
	return response.BookResponse(book)
}

// Update implements BookService
func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	errorhelper.PanicIfErr(err)
	book.BookName = request.BookName
	b.BookRepository.Update(ctx, book)
}