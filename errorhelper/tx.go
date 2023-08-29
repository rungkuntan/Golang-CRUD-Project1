package errorhelper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
  err := recover()

  if err != nil {
	errRollBack := tx.Rollback()
	PanicIfErr(errRollBack)
	panic(err)
  } else {
	errCommit := tx.Commit()
	PanicIfErr(errCommit)
  }
}