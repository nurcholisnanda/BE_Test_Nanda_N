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

func TestQueryItems(t *testing.T) {
	t.Run("Should query items correctly", func(t *testing.T) {
		resolvers := Resolver{
			items:        ItemTest,
			checkoutItem: &model.CheckoutItem{},
		}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		var resp struct {
			items struct {
				items []struct {
					Sku, Name string
					Price     float64
					Qty       int
				}
			}
		}
		q := `
			query findItems {
				items {
				SKU
				Name
				Price
				Qty
				}
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, err)
	})
}

func TestCheckout(t *testing.T) {
	item := &model.Item{
		Sku:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}
	ItemTest = append(ItemTest, item)
	item = &model.Item{
		Sku:   "43N23P",
		Name:  "MacBook Pro",
		Price: 5399.99,
		Qty:   5,
	}
	ItemTest = append(ItemTest, item)
	item = &model.Item{
		Sku:   "A304SD",
		Name:  "Alexa Speaker",
		Price: 109.5,
		Qty:   10,
	}
	ItemTest = append(ItemTest, item)
	item = &model.Item{
		Sku:   "234234",
		Name:  "Raspberry Pi B",
		Price: 30,
		Qty:   2,
	}
	ItemTest = append(ItemTest, item)

	t.Run("Test Checkout All Item First Combination", func(t *testing.T) {
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
				Checkout(input: [{Name:"Google Home",Qty:9},{Name:"MacBook Pro",Qty:1},{Name:"Alexa Speaker",Qty:3},{Name:"Alexa Speaker",Qty:4},{Name:"Raspberry Pi B",Qty:2}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, err)
	})

	t.Run("Test Checkout All Item Second Combination", func(t *testing.T) {
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
		err := c.Post(q, &resp)
		require.Nil(t, err)
	})

	t.Run("Test Checkout RaspBerry Pi B", func(t *testing.T) {
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
				Checkout(input: [{Name:"Raspberry Pi B",Qty:1}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, err)
	})

	t.Run("Test Checkout RaspBerry Pi B more than Qty", func(t *testing.T) {
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
				Checkout(input: [{Name:"Raspberry Pi B",Qty:3}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, resp.Checkout.items)
		require.Equal(t, float64(0), resp.Checkout.Price)
		require.NotNil(t, err)
	})

	t.Run("Test Checkout Alexa Speaker more than Qty", func(t *testing.T) {
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
				Checkout(input: [{Name:"Alexa Speaker",Qty:12}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, resp.Checkout.items)
		require.Equal(t, float64(0), resp.Checkout.Price)
		require.NotNil(t, err)
	})

	t.Run("Test Checkout MacBook Pro more than Qty", func(t *testing.T) {
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
				Checkout(input: [{Name:"MacBook Pro",Qty:12}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, resp.Checkout.items)
		require.Equal(t, float64(0), resp.Checkout.Price)
		require.NotNil(t, err)
	})

	t.Run("Test Checkout Google Home more than Qty", func(t *testing.T) {
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
				Checkout(input: [{Name:"Google Home",Qty:12}]) {
					items {
					  SKU
					  Name
					  Qty
					}
					Price
				  }
			}
		`
		err := c.Post(q, &resp)
		require.Nil(t, resp.Checkout.items)
		require.Equal(t, float64(0), resp.Checkout.Price)
		require.NotNil(t, err)
	})
}
