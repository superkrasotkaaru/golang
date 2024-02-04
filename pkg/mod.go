package aru

import (
	"errors"
)

type Wish struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

var wishlist = []Wish{
	{
		Id:      "1",
		Title:   "Cartier",
		Type:    "Gold",
	},
	{
		Id:      "2",
		Title:   "Airpods max",
		Type:    "Blue",
	},
	{
		Id:      "3",
		Title:   "Camera",
		Type:    "Canon G7X",
	},
	{
		Id:      "4",
		Title:   "Books",
		Type:    "Pride and Prejudice, Little Women, Save me",
	},
	{
		Id:      "5",
		Title:   "Villa",
		Type:    "in Maldives",
	},
}

func GetWishlist() []Wish {
	return wishlist
}

func GetWish(id string) (*Wish, error) {
	for _, r := range wishlist {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Wish not found")
}