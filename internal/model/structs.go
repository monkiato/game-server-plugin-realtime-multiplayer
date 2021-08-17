package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	gorm.Model
	MaxCapacity uint
	CanJoin bool //match can be locked even if not at max capacity
	Visibility uint //friends-only, public, private, etc
	EndedAt time.Time
}

type Player struct {
	gorm.Model
	socialID string
	socialNetwork string
}

type MatchPlayer struct {
	gorm.Model
	MatchID uint
	Match Match
	PlayerID uint
	Player Player
}
