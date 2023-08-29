package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_project/errorhelper"
	"go_project/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx,err := b.Db.Begin()
	errorhelper.PanicIfErr(err)
	defer errorhelper.CommitOrRollback(tx)

	SQL := "delete from book where id =$1"
	_,errQuery := tx.ExecContext(ctx,SQL,bookId)
	errorhelper.PanicIfErr(errQuery)
}


func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book  {
    tx,err := b.Db.Begin()
	errorhelper.PanicIfErr(err)
	defer errorhelper.CommitOrRollback(tx)

	SQL := "select id,name from book"
	result,errexec := tx.QueryContext(ctx,SQL)
	errorhelper.PanicIfErr(errexec)
	defer result.Close()

    var books []model.Book

	for result.Next() {
		book := model.Book{}
        err := result.Scan(&book.Id,&book.Name)
		errorhelper.PanicIfErr(err)
		books = append(books,book)
	}
	return books
}


func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	tx,err := b.Db.Begin()
	errorhelper.PanicIfErr(err)
    defer errorhelper.CommitOrRollback(tx)
	
	SQL := " select id,name from book where id =$1"
	result,errQuery := tx.QueryContext(ctx,SQL,bookId)
	errorhelper.PanicIfErr(errQuery)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id,&book.Name)
		errorhelper.PanicIfErr(err)
		return book,nil
	}else {
       return book, errors.New("Book id not found")
	}
	
}

func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
  tx,err := b.Db.Begin()
  errorhelper.PanicIfErr(err)
  defer errorhelper.CommitOrRollback(tx)

  SQL :=  "insert into book(name) values($1)"
  _,errExec := tx.ExecContext(ctx,SQL,book.Name)
  errorhelper.PanicIfErr(errExec)

}


func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
  tx,err := b.Db.Begin()
  errorhelper.PanicIfErr(err)
  defer errorhelper.CommitOrRollback(tx)

  SQL := "update book set name=$1 where id=$2"
  _,errExec := tx.ExecContext(ctx,SQL,book.Name,book.Id)
  errorhelper.PanicIfErr(errExec)
}