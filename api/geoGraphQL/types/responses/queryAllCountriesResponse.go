package responses

import "github.com/reoverduguez/map-app/backend/api/geoGraphQL/types/shared"

type QueryAllCountriesResponse struct {
	Data struct {
		Countries Countries `json:"countries"`
	}
}

type Countries struct {
	TotalCount int             `json:"totalCount"`
	Edges      []Edge          `json:"edges"`
	PageInfo   shared.PageInfo `json:"pageInfo"`
}

type Edge struct {
	Cursor string `json:"cursor"`
	Node   Node   `json:"node"`
}

type Node struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Iso3  string `json:"iso3"`
	Emoji string `json:"emoji"`
}
