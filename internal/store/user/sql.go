package user

import (
	"log"
	"time"
	"errors"
	"github.com/NegarMov/shopping-api/internal/model"
	"gorm.io/gorm"
)

type SQLItem struct {
	ID      	uint				`gorm:"primaryKey;autoIncrement"`
	Username    string				`gorm:"not null"`
	Password    string				`gorm:"not null"`
	CreatedAt	time.Time			`gorm:"not null"`
}

func (SQLItem) TableName() string {
	return "user"
}

type SQL struct {
	DB *gorm.DB
}

func NewSQL(db *gorm.DB) User {
	if err := db.AutoMigrate(new(SQLItem)); err != nil {
		log.Fatal(err)
	}

	return SQL{
		DB: db,
	}
}

func (sql SQL) SignUp(u model.User) (model.User, error) {
	var existingUser SQLItem
	
	if err := sql.DB.Where("username = ?", u.Username).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, err
		}
	}

	if existingUser.ID != 0 {
		return model.User{}, ErrDuplicatedUsername
	}
	
	item := SQLItem{
		Username: 	u.Username,
		Password: 	u.Password,
		CreatedAt:	u.CreatedAt,
	}

	if err := sql.DB.Create(&item).Error; err != nil {
		return model.User{}, err
	}
	
	return model.User{
		ID:   		item.ID,
		Username:	item.Username,
		Password:	item.Password,
		CreatedAt:	item.CreatedAt,
	}, nil
}

func (sql SQL) SignIn(username string, password string) (model.User, error) {
	var item SQLItem
	
	err := sql.DB.Where("username = ? AND password = ?", username, password).First(&item).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, ErrWrongCredentials
		}
		return model.User{}, err
	}
	
	return model.User{
		ID:   		item.ID,
		Username:	item.Username,
		Password:	item.Password,
		CreatedAt:	item.CreatedAt,
	}, nil
}
