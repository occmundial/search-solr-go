package solr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/solr-go"
)

func TestQuery(t *testing.T) {
	a := assert.New(t)
	got, err := solr.NewQuery().
		QueryParser(
			solr.NewDisMaxQueryParser().
				Query("'solr rocks'"),
		).
		Queries(solr.M{
			"query_filters": []solr.M{
				{
					"#size_tag": solr.M{
						"field": solr.M{
							"f":     "size",
							"query": "XL",
						},
					},
				},
				{
					"#color_tag": solr.M{
						"field": solr.M{
							"f":     "color",
							"query": "Red",
						},
					},
				},
			},
		}).
		Facets(
			solr.NewTermsFacet("categories").
				Field("cat").Limit(10),
			solr.NewQueryFacet("high_popularity").
				Query("popularity:[8 TO 10]"),
		).
		Sort("score").
		Offset(1).
		Limit(10).
		Filters("inStock:true").
		Fields("name", "price").
		BuildJSON()
	a.NoError(err)
	expect := `{"facet":{"categories":{"field":"cat","limit":10,"type":"terms"},"high_popularity":{"q":"popularity:[8 TO 10]","type":"query"}},"fields":["name","price"],"filter":["inStock:true"],"limit":10,"offset":1,"queries":{"query_filters":[{"#size_tag":{"field":{"f":"size","query":"XL"}}},{"#color_tag":{"field":{"f":"color","query":"Red"}}}]},"query":"{!dismax v='solr rocks'}","sort":"score"}`
	a.Equal(expect, string(got))
}
