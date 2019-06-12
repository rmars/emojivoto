package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name     string
	FavEmoji string
	PublicId string
}

type DBClient struct {
	db *gorm.DB
}

func NewClient() (*DBClient, error) {
	db, err := gorm.Open("sqlite3", "temp.db")
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// put some fake data in our db
	seedDatabase(db)

	return &DBClient{db}, nil
}

func (d *DBClient) GetUsers(where ...interface{}) ([]*User, error) {
	users := []*User{}
	return users, d.db.Find(&users, where...).Error
}

func seedDatabase(db *gorm.DB) {
	// Clear previous db data
	db.Delete(&User{})

	// Create some fake users
	db.Create(&User{Name: "Rey", FavEmoji: ":100:", PublicId: "cGVvcGxlOjg1"})
	db.Create(&User{Name: "BB8", FavEmoji: ":fire:", PublicId: "cGVvcGxlOjg3"})
	db.Create(&User{Name: "Captain Phasma", FavEmoji: ":doughnut:", PublicId: "cGVvcGxlOjg4"})
	db.Create(&User{Name: "R2-D2", FavEmoji: "fire", PublicId: "cGVvcGxlOjM="})
	db.Create(&User{Name: "Leia Organa", FavEmoji: ":champagne:", PublicId: "cGVvcGxlOjU="})
	db.Create(&User{Name: "Padmé Amidala", FavEmoji: ":cat2:", PublicId: "cGVvcGxlOjM1"})
	db.Create(&User{Name: "Jocasta Nu", FavEmoji: ":dog:", PublicId: "cGVvcGxlOjc0"})
	db.Create(&User{Name: "C-3PO", FavEmoji: ":pray:", PublicId: "c3BlY2llczoy"})
}
