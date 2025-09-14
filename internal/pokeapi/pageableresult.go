package pokeapi

type PageableResult[T any] struct {
	Count       int    `json:"count"`
	NextUrl     string `json:"next"`
	PreviousUrl string `json:"previous"`
	Items       []T    `json:"results"`
}
