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
	Cards    []Card  `gorm:"foreignKey:UserID"` // One-to-Many relationship with Card
	Tags     []Tag   `gorm:"foreignKey:UserID"` // One-to-Many relationship with Tag
	States   []State `gorm:"foreignKey:UserID"` // One-to-Many relationship with State
}

type Card struct {
	ID          string `gorm:"type:char(36);primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(60);not null" json:"title" binding:"required"`
	Description string `gorm:"type:text;not null" json:"description"`
	Position    int    `gorm:"not null" json:"position" binding:"required"`
	StateID     int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"state_id" binding:"required"`
	UserID      string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User        User   `gorm:"foreignKey:UserID" binding:"-" json:"-"`
	State       State  `gorm:"foreignKey:StateID" binding:"-" json:"-"`
	Tags        []Tag  `gorm:"many2many:card_tags;foreignKey:ID;references:ID" json:"tags"`
}

type Tag struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20);unique" json:"name" binding:"required"`
	Color  string `gorm:"type:varchar(6);not null" json:"color" binding:"required"`
	UserID string `gorm:"type:char(36);not null" json:"-"`
	User   User   `gorm:"foreignKey:UserID" binding:"-" json:"-"`
	Cards  []Card `gorm:"many2many:card_tags;foreignKey:ID;references:ID" json:"-"`
}

type State struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
	Position int    `gorm:"not null" json:"position" binding:"required"`
	Color    string `gorm:"type:char(6);not null" json:"color" binding:"required"`
	UserID   string `gorm:"type:char(36);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User     User   `gorm:"foreignKey:UserID" binding:"-" json:"-"`
	Cards    []Card `gorm:"foreignKey:StateID" json:"-"` // One-to-Many relationship with Card
}

var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./data/minban.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Card{}, &Tag{}, &State{}); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}
