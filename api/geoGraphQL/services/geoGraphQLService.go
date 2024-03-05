package services

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/reoverduguez/map-app/backend/api/geoGraphQL/types/responses"
)

var baseURL = "https://api.geographql.rudio.dev/graphql"

func QueryAllCountiresByRegion(region *string, subregion *string) (responses.QueryAllCountriesResponse, error) {
	var filter string
	var response responses.QueryAllCountriesResponse

	client := graphql.NewClient(baseURL)

	if region != nil && subregion != nil {
		filter = fmt.Sprintf(`{ region: %s, subregion: %s }`, *region, *subregion)
	} else if region != nil {
		filter = fmt.Sprintf(`{ region: %s }`, *region)
	} else if subregion != nil {
		filter = fmt.Sprintf(`{ subregion: %s }`, *subregion)
	} else {
		filter = "{}"
	}

	query := fmt.Sprintf(`{
		countries(
			filter: %s
			page: { first: 10, after: "eyJjdXJzb3IiOjE1Mn0" }
		) {
			totalCount
			edges {
			cursor
			node {
				id
				name
				iso3
				emoji
			}
			}
			pageInfo {
				hasNextPage
				hasPreviousPage
				endCursor
				startCursor
			}
		}
	  }`, filter)

	request := graphql.NewRequest(query)
	err := client.Run(context.Background(), request, &response)

	if err != nil {
		return responses.QueryAllCountriesResponse{}, err
	}

	return response, nil
}
