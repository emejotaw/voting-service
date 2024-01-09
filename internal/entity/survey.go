package entity

type Survey struct {
	ID          string
	Title       string
	Description string
	Option      *[]Option
}
