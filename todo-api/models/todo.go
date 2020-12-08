package models

// Todo r
type Todo struct {
	ID			int			`json:"id"`
	Name		string		`json:"name" pg:",notnull"`
	Done		bool		`json:"done"`
}