package model

const defaultLimit = 10

type SearchRequest struct {
	Status  string
	Offset  uint
	Limit   uint
	OwnerId string
}

type Paging struct {
	Total  uint `json:"total"`
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
}

type SearchResponse struct {
	Paging  Paging `json:"paging"`
	Results []Pet  `json:"results"`
}

func NewSearchRequest() SearchRequest {
	return SearchRequest{Offset: 0, Limit: defaultLimit}
}
