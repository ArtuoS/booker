package service

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ArtuoS/booker-api/internal/entity"
	"github.com/ArtuoS/booker-api/internal/infra"
)

type CsvReader struct {
	Path string
	Db   *sql.DB
}

func NewCsvReader(path string, db *sql.DB) CsvReader {
	return CsvReader{
		Path: path,
		Db:   db,
	}
}

func (c *CsvReader) Start() error {
	file, err := os.Open("../assets/files/authors.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	authorService := NewAuthorService(infra.NewAuthorRepository(c.Db))

	start := time.Now()

	authors, lines, done := make(chan entity.Author), make(chan string), make(chan bool)
	workers := 25
	for worker := 0; worker < workers; worker++ {
		go func() {
			for author := range authors {
				lines <- author.Name
			}

			done <- true
		}()
	}

	go readCsv(file, authors)

	waits := workers
	defaultCmd := "INSERT INTO authors (name) VALUES "
	cmd := defaultCmd
	rows := 0
	for {
		select {
		case line := <-lines:
			rows++
			cmd += fmt.Sprintf("('%s'),", line)
			if rows == 10000 {
				cmd = cmd[0 : len(cmd)-1]
				authorService.Execute(cmd)
				cmd = defaultCmd
				rows = 0
			}
		case <-done:
			waits--
			if waits == 0 {
				if rows > 0 {
					cmd = cmd[0 : len(cmd)-1]
					authorService.Execute(cmd)
				}

				elapsed := time.Since(start)
				fmt.Println("Time elapsed: ", elapsed)
				return nil
			}
		}
	}
}

func readCsv(b io.Reader, ch chan entity.Author) {
	r := csv.NewReader(b)
	defer close(ch)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		author := entity.NewAuthor(record[0])
		ch <- author
	}
}
