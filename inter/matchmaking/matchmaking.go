package matchmaking

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
)

var Games = make(map[string]*network.Client)

func FindGIDs() string {
	var gameID string

	for z := range Games {
		gameID = z
	}
	return gameID
}
