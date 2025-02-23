package handlers

import "github.com/Arismonx/Golang-GraphQl/graph/model"

var players = []*model.Player{
	{
		ID:    "PLAYER-456",
		Name:  "Tuschy",
		Leval: 9999,
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
		Leval: 1000000,
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

func CreatePlayer(req *model.NewPlayer) *model.Player {
	newPlayer := &model.Player{
		ID:    "PLAYER-001",
		Name:  "Himari",
		Leval: 100,
		Class: &req.Class,
		Item: []*model.Item{
			{
				ID:   "ITEM 003",
				Name: "Sword",
			},
		},
	}

	players = append(players, newPlayer)
	return newPlayer
}
