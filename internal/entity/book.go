package entity

type Book struct {
	Id              int64    `json:"id"`
	Name            string   `json:"name"`
	Edition         string   `json:"edition"`
	PublicationYear int      `json:"publication_year"`
	Authors         []Author `json:"authors"`
}
