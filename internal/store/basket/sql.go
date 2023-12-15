package basket

import (
	"log"
	"time"
	"errors"
	"github.com/jackc/pgtype"
	"github.com/NegarMov/shopping-api/internal/model"
	"gorm.io/gorm"
)

type SQLItem struct {
	ID      	uint			`gorm:"primaryKey;autoIncrement"`
	CreatedAt	time.Time		`gorm:"not null"`
	UpdatedAt	time.Time 		`gorm:"not null"`
	Data		pgtype.JSONB	`gorm:"type:jsonb;default:'{}'"`
	State		model.State		`gorm:"not null"`
	UserID		uint			`gorm:"not null;many2one:basket_user"`
}

func (SQLItem) TableName() string {
	return "basket"
}

type SQL struct {
	DB *gorm.DB
}

func NewSQL(db *gorm.DB) Basket {
	if err := db.AutoMigrate(new(SQLItem)); err != nil {
		log.Fatal(err)
	}

	return SQL{
		DB: db,
	}
}

func (sql SQL) GetAll() ([]model.Basket, error) {
	var items []SQLItem

	if err := sql.DB.Model(new(SQLItem)).Find(&items).Error; err != nil {
		return nil, err
	}

	baskets := make([]model.Basket, 0)

	for _, item := range items {
		baskets = append(baskets, model.Basket{
			ID:   		item.ID,
			CreatedAt:	item.CreatedAt,
			UpdatedAt:	item.UpdatedAt,
			Data:		item.Data,
			State:		item.State,
			UserID:		item.UserID,
		})
	}

	return baskets, nil
}

func (sql SQL) Create(b model.Basket) (model.Basket, error) {
	item := SQLItem{
		CreatedAt:	b.CreatedAt,
		UpdatedAt:	b.UpdatedAt,
		Data:		b.Data,
		State:		b.State,
		UserID:		b.UserID,
	}

	if err := sql.DB.Create(&item).Error; err != nil {
		return model.Basket{}, err
	}
	
	return model.Basket{
		ID:   		item.ID,
		CreatedAt:	item.CreatedAt,
		UpdatedAt:	item.UpdatedAt,
		Data:		item.Data,
		State:		item.State,
		UserID:		item.UserID,
	}, nil
}

func (sql SQL) Update(id uint, new_b model.Basket) (model.Basket, error) {
	var b SQLItem

	if err := sql.DB.First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Basket{}, ErrBasketNotFound
		}

		return model.Basket{}, err
	}

	if b.State == model.Completed {
		return model.Basket{}, ErrBasketCompleted
	}

	if err := sql.DB.Model(&b).Updates(&SQLItem{
		UpdatedAt:	new_b.UpdatedAt,
		Data:		new_b.Data,
		State:		new_b.State,
	}).Error; err != nil {
		return model.Basket{}, err
	}

	return model.Basket{
		ID:   		b.ID,
		CreatedAt:	b.CreatedAt,
		UpdatedAt:	b.UpdatedAt,
		Data:		b.Data,
		State:		b.State,
		UserID:		b.UserID,
	}, nil
}

func (sql SQL) Get(id uint) (model.Basket, error) {
	var b SQLItem

	if err := sql.DB.First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Basket{}, ErrBasketNotFound
		}

		return model.Basket{}, err
	}

	return model.Basket{
		ID:   		b.ID,
		CreatedAt:	b.CreatedAt,
		UpdatedAt:	b.UpdatedAt,
		Data:		b.Data,
		State:		b.State,
		UserID:		b.UserID,
	}, nil
}

func (sql SQL) Delete(id uint) (model.Basket, error) {
	var b SQLItem

	if err := sql.DB.First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Basket{}, ErrBasketNotFound
		}

		return model.Basket{}, err
	}

	if err := sql.DB.Delete(&b).Error; err != nil {
		return model.Basket{}, err
	}

	return model.Basket{
		ID:   		b.ID,
		CreatedAt:	b.CreatedAt,
		UpdatedAt:	b.UpdatedAt,
		Data:		b.Data,
		State:		b.State,
		UserID:		b.UserID,
	}, nil
}
