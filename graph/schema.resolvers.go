package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BE_TEST/graph/generated"
	"BE_TEST/graph/model"
	"context"
	"errors"
)

var items []*model.Item

func init() {
	item := &model.Item{
		Sku:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}
	items = append(items, item)
	item = &model.Item{
		Sku:   "43N23P",
		Name:  "MacBook Pro",
		Price: 5399.99,
		Qty:   5,
	}
	items = append(items, item)
	item = &model.Item{
		Sku:   "A304SD",
		Name:  "Alexa Speaker",
		Price: 109.5,
		Qty:   10,
	}
	items = append(items, item)
	item = &model.Item{
		Sku:   "234234",
		Name:  "Raspberry Pi B",
		Price: 30,
		Qty:   2,
	}
	items = append(items, item)
}

func (r *mutationResolver) Checkout(ctx context.Context, input []*model.InputItem) (*model.CheckoutItem, error) {
	var inputItems []*model.CheckItem
	var tmpPrice float64
	var hasMac bool

	for i := 0; i < len(input); i++ {
		if input[i].Name == "MacBook Pro" {
			hasMac = true
		}
	}

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i].Name == input[j].Name {
				input[i].Qty = input[i].Qty + input[j].Qty
				input[j].Qty = 0
			}
		}
	}

	for i := 0; i < len(input); i++ {
		inputItem := &model.CheckItem{
			Name: input[i].Name,
			Qty:  input[i].Qty,
		}
		for j := 0; j < len(items); j++ {
			if inputItem.Qty == 0 {
				break
			}
			if inputItem.Name == items[j].Name {
				if inputItem.Name == "Google Home" {
					if inputItem.Qty > items[j].Qty {
						return nil, errors.New("Sorry " + inputItem.Name + " Qty exceeding our stock, couldn't continue the checkout")
					}
					inputItem.Sku = items[j].Sku
					tmpQty := inputItem.Qty / 3
					tmpModQty := inputItem.Qty % 3
					tmpPrice = tmpPrice + (float64(tmpQty) * 2 * items[j].Price) + (float64(tmpModQty) * items[j].Price)
					inputItems = append(inputItems, inputItem)
				} else if inputItem.Name == "Alexa Speaker" {
					if inputItem.Qty > items[j].Qty {
						return nil, errors.New("Sorry " + inputItem.Name + " Qty exceeding our stock, couldn't continue the checkout")
					}
					inputItem.Sku = items[j].Sku
					if inputItem.Qty > 3 {
						tmpPrice = tmpPrice + (0.9 * items[j].Price * float64(inputItem.Qty))
					}
					inputItems = append(inputItems, inputItem)
				} else if inputItem.Name == "MacBook Pro" {
					if inputItem.Qty > items[j].Qty {
						return nil, errors.New("Sorry " + inputItem.Name + " Qty exceeding our stock, couldn't continue the checkout")
					}
					inputItem.Sku = items[j].Sku
					tmpPrice = tmpPrice + (items[j].Price * float64(inputItem.Qty))
					inputItems = append(inputItems, inputItem)
					newInputItem := &model.CheckItem{
						Sku:  "234234",
						Name: "Raspberry Pi B",
					}
					if inputItem.Qty >= 2 {
						newInputItem.Qty = 2
					} else {
						newInputItem.Qty = 1
					}
					inputItems = append(inputItems, newInputItem)
				} else if inputItem.Name == "Raspberry Pi B" {
					if hasMac {
						break
					}
					if inputItem.Qty > items[j].Qty {
						return nil, errors.New("Sorry " + inputItem.Name + " Qty exceeding our stock, couldn't continue the checkout")
					}
					inputItem.Sku = items[j].Sku
					tmpPrice = tmpPrice + (items[j].Price * float64(inputItem.Qty))
					inputItems = append(inputItems, inputItem)
				} else {
					return nil, errors.New("Sorry the item you request is unavailable")
				}
				break
			}
		}
	}
	checkoutItems := &model.CheckoutItem{
		Items: inputItems,
		Price: tmpPrice,
	}
	r.checkoutItem = checkoutItems
	return r.checkoutItem, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	r.items = items
	return r.items, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
