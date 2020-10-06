package graph

import "BE_TEST/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	items        []*model.Item
	checkoutItem *model.CheckoutItem
}
