package graph

import (
	"BE_TEST/graph/generated"
	"BE_TEST/graph/model"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

var ItemTest []*model.Item

func TestCheckout(t *testing.T) {
	t.Run("Test Checkout All Item", func(t *testing.T) {
		item := &model.Item{
			Sku:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Qty:   10,
		}
		ItemTest = append(items, item)
		item = &model.Item{
			Sku:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99,
			Qty:   5,
		}
		ItemTest = append(items, item)
		item = &model.Item{
			Sku:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.5,
			Qty:   10,
		}
		ItemTest = append(items, item)
		item = &model.Item{
			Sku:   "234234",
			Name:  "Raspberry Pi B",
			Price: 30,
			Qty:   2,
		}
		ItemTest = append(items, item)
		resolvers := Resolver{
			items:        ItemTest,
			checkoutItem: &model.CheckoutItem{},
		}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		var resp struct {
			Checkout struct {
				items []struct {
					Sku, Name string
					Qty       int
				}
				Price float64
			}
		}
		q := `
			mutation { 
				Checkout(input: [{Name:"Google Home",Qty:9},{Name:"MacBook Pro",Qty:2},{Name:"Alexa Speaker",Qty:3},{Name:"Alexa Speaker",Qty:4},{Name:"Raspberry Pi B",Qty:1}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		c.MustPost(q, &resp)
		require.NotNil(t, &resp)
	})
}
