package models

type SignRequest struct {
	Name     string
	Password string
}

type UpdateRequest struct {
	Alias string
	Jwt   string
}

type ShortRequest struct {
	Jwt string
	Url string
}

type GetRequest struct {
	Url string
}
