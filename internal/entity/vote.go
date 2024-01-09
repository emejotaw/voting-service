package entity

type Vote struct {
	ID        string
	IpAddress string
	UserAgent string
	Email     string
	OptionID  string
	Option    *Option
}
