package game

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/internal/model"
	"github.com/sirupsen/logrus"
	"time"
)

type PlayerConnection struct {

}

type MatchRegistry struct {
	Id uint
	Connections []*PlayerConnection
}

func NewMatchRegistry(matchId uint) *MatchRegistry {
	return &MatchRegistry{
		Id:          matchId,
		Connections: []*PlayerConnection{},
	}
}

type MatchInfo struct {
	Capacity uint
	Visibility uint
}

func CreateMatch(db *gorm.DB, playerToken string, matchInfo *MatchInfo) (*model.Match, error) {

	//TODO: implement WHERE token = {playerToken}, but we need a credential/user-login plugin first for identity related stuff
	//fetch player data from DB
	//var player model.Player
	//if err := db.First(&player, playerId).Error; err != nil {
	//	returnError := fmt.Errorf("unable to find player %s", playerId)
	//	logrus.Errorf("%s\n%v", returnError.Error(), err.Error())
	//	return nil, returnError
	//}

	//TODO: remove, this is a hardcoded user
	player := &model.Player{
		Model: gorm.Model{
			ID:        1234,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
	}

	newMatch := &model.Match{
		MaxCapacity:     matchInfo.Capacity,
		CanJoin:         true,
		Visibility:      matchInfo.Visibility,
	}
	if err := db.Create(newMatch).Error; err != nil {
		returnError := fmt.Errorf("unable to create new match")
		logrus.Errorf("%s\n%v", returnError.Error(), err.Error())
		return nil, returnError
	}

	matchPlayer := &model.MatchPlayer{
		MatchID:  newMatch.ID,
		PlayerID: player.ID,
	}
	if err := db.Create(matchPlayer).Error; err != nil {
		returnError := fmt.Errorf("unable to add player %d to match %d", player.ID, newMatch.ID)
		logrus.Errorf("%s\n%v", returnError.Error(), err.Error())
		return nil, returnError
	}

	return newMatch, nil
}