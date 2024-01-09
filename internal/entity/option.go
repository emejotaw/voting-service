package entity

type Option struct {
	ID          string
	Title       string
	Description string
	SurveyID    string
	Survey      *Survey
}
