package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       string  `gorm:"type:char(36);primaryKey"`
	Name     string  `gorm:"type:varchar(40);not null;unique" json:"name" binding:"required"`
	Password string  `gorm:"type:char(64);not null" json:"password" binding:"required"`
	Boards   []Board `gorm:"foreignKey:UserID" json:"boards"` // One-to-Many relationship with Board
}

type Board struct {
	ID          string  `gorm:"type:char(36);primaryKey" json:"id"`
	Title       string  `gorm:"type:varchar(40);not null;unique" json:"name" binding:"required"`
	Description string  `gorm:"type:text;not null" json:"description"`
	Token       string  `gorm:"type:char(4);not null;unique" json:"token" binding:"required"`
	Selected    bool    `gorm:"default:false" json:"selected"`
	UserID      string  `gorm:"type:char(36);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User        User    `gorm:"foreignKey:UserID" binding:"-" json:"-"`
	Cards       []Card  `gorm:"foreignKey:BoardID" json:"cards"`  // One-to-Many relationship with Card
	Tags        []Tag   `gorm:"foreignKey:BoardID" json:"tags"`   // One-to-Many relationship with Tag
	States      []State `gorm:"foreignKey:BoardID" json:"states"` // One-to-Many relationship with State
}

type Card struct {
	ID          int    `gorm:"primaryKey;autoIncrement;autoIncrementStart:1" json:"id"`
	Title       string `gorm:"type:varchar(60);not null" json:"title" binding:"required"`
	Description string `gorm:"type:text;not null" json:"description"`
	Position    int    `gorm:"not null" json:"position" binding:"required"`
	StateID     int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"state_id" binding:"required"`
	BoardID     string `gorm:"type:char(36);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Board       Board  `gorm:"foreignKey:BoardID" binding:"-" json:"-"`
	State       State  `gorm:"foreignKey:StateID" binding:"-" json:"-"`
	Tags        []Tag  `gorm:"many2many:card_tags;" json:"tags"`
}

type Tag struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(20);unique" json:"name" binding:"required"`
	Color   string `gorm:"type:varchar(6);not null" json:"color" binding:"required"`
	BoardID string `gorm:"type:char(36);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Board   Board  `gorm:"foreignKey:BoardID" json:"-"`
	Cards   []Card `gorm:"many2many:card_tags;" json:"-"`
}

type State struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
	Position int    `gorm:"not null" json:"position" binding:"required"`
	Color    string `gorm:"type:char(6);not null" json:"color" binding:"required"`
	BoardID  string `gorm:"type:char(36);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Board    Board  `gorm:"foreignKey:BoardID" json:"-"`
	Cards    []Card `gorm:"foreignKey:StateID" json:"-"` // One-to-Many relationship with Card
}

var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./data/minban.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Board{}, &Card{}, &Tag{}, &State{}); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}
