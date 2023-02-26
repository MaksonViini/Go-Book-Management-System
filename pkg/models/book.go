package models

import (
	"github.com/maksonviini/Go-Book-Management-System/pkg/config"
)

type Book struct {
	Id          int64  `json:"id"`
	Name        string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	// createdAt time.Time `json:"createdAt"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// }

func Insert(book Book) (id int64, err error) {
	conn, err := config.Connect()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO Books (Name, Author, Publication) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, book.Name, book.Author, book.Publication).Scan(&id)

	return
}

func Get(id int64) (book Book, err error) {
	conn, err := config.Connect()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM Book WHERE id=$1`, id)

	err = row.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)

	return
}

func GetAll() (books []Book, err error) {
	conn, err := config.Connect()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM Book`)

	if err != nil {
		return
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)

		if err != nil {
			continue
		}
		books = append(books, book)
	}

	return
}

func Update(id int64, book Book) (int64, error) {
	conn, err := config.Connect()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE book SET Name=$2, Author=$3, Publication=$4 WHERE Id=$1`, id, book.Name, book.Author, book.Publication)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {
	conn, err := config.Connect()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM Book WHERE Id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
