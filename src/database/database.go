package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:char(36);primaryKey"`
	Name     string `gorm:"type:varchar(40);not null" json:"name" binding:"required"`
	Password string `gorm:"type:char(64);not null" json:"password" binding:"required"`
	Cards    []Card `gorm:"foreignKey:UserID"` // One-to-Many relationship with Card
}

type Card struct {
	ID          string `gorm:"type:char(36);primaryKey"`
	Title       string `gorm:"type:varchar(60);not null" json:"title" binding:"required"`
	Description string `gorm:"type:text;not null" json:"description"`
	Position    int    `gorm:"not null" json:"position" binding:"required"`
	StateID     int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"state_id" binding:"required"`
	UserID      string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Tags        []Tag  `gorm:"many2many:card_tags;joinForeignKey:CardID;joinReferences:TagName" json:"tags"`
}

type Tag struct {
	Name  string `gorm:"type:varchar(20);primaryKey" json:"name" binding:"required"`
	Color string `gorm:"type:varchar(6);not null" json:"color" binding:"required"`
	Cards []Card `gorm:"many2many:card_tags;joinForeignKey:TagName;joinReferences:CardID" json:"-"`
}

type State struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
	Position int    `gorm:"not null;unique" json:"position" binding:"required"`
	Color    string `gorm:"type:char(6);not null" json:"color" binding:"required"`
	Cards    []Card `gorm:"foreignKey:StateID" json:"-"` // One-to-Many relationship with Card
}

type ActivityLog struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	CardID    string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Timestamp string `gorm:"type:datetime;not null"`
	ActionID  int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ActivityAction struct {
	ID         int           `gorm:"primaryKey"`
	ActionName string        `gorm:"type:varchar(20);not null"`
	Logs       []ActivityLog `gorm:"foreignKey:ActionID"` // One-to-Many relationship with ActivityLog
}

var DB *gorm.DB

func InitializeDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./data/minban.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Card{}, &Tag{}, &State{}, &ActivityLog{}, &ActivityAction{}); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	InitializeStates()
	InitializeTags()
}

func InitializeStates() {
	states := []State{
		{ID: 1, Name: "Todo", Position: 1, Color: "FF0000"},
		{ID: 2, Name: "In Progress", Position: 2, Color: "00FF00"},
		{ID: 3, Name: "Done", Position: 3, Color: "0000FF"},
	}

	for _, state := range states {
		if err := DB.Where("id = ?", state.ID).FirstOrCreate(&state).Error; err != nil {
			log.Fatalf("Failed to create state: %v", err)
		}
	}
}

func InitializeTags() {
	tags := []Tag{
		{Name: "Feature", Color: "FF0000"},
		{Name: "Bug", Color: "00FF00"},
	}

	for _, tag := range tags {
		if err := DB.Where("name = ?", tag.Name).FirstOrCreate(&tag).Error; err != nil {
			log.Fatalf("Failed to create tag: %v", err)
		}
	}
}
