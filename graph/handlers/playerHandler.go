package handlers

import "github.com/Arismonx/Golang-GraphQl/graph/model"

var players = []*model.Player{
	{
		ID:    "PLAYER-456",
		Name:  "Tuschy",
		Level: 9999,
		Class: &model.AllPlayerClass[2],
		Item: []*model.Item{
			{
				ID:   "ITEM 001",
				Name: "BOW",
			},
		},
	},
	{
		ID:    "PLAYER-111",
		Name:  "Lotus",
		Level: 1000000,
		Class: &model.AllPlayerClass[1],
		Item: []*model.Item{
			{
				ID:   "ITEM 002",
				Name: "SAFFT",
			},
		},
	},
}

func GetPlayers() []*model.Player {
	return players
}
