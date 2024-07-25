package repository

import (
	"smart-kost-backend/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(inputModel *model.UserList) (*model.UserList, error)
	GetUserByUsername(inputModel *model.UserList) (*model.UserList, error)
	UpdateIsOnline(inputModel *model.UserList) (*model.UserList, error)
	UpdateIsSOS(inputModel *model.UserList) (*model.UserList, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) WithTx(tx *gorm.DB) UserRepository {
	return &userRepo{
		db: tx,
	}
}

func (r *userRepo) CreateUser(inputModel *model.UserList) (*model.UserList, error) {

	if dbc := r.db.Create(inputModel).Scan(inputModel); dbc.Error != nil {
		return nil, dbc.Error
	}

	return inputModel, nil
}

func (r *userRepo) GetUserByUsername(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Where("username = ?", inputModel.Username).First(&model.UserList{}).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r *userRepo) UpdateIsOnline(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Raw("UPDATE user_list SET is_online = ?, updated_at = ? WHERE user_id = ? RETURNING *", string(inputModel.IsOnline), time.Now(), inputModel.UserId).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}

func (r *userRepo) UpdateIsSOS(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Raw("UPDATE user_list SET is_sos = ?, updated_at = ? WHERE user_id = ? RETURNING *", string(inputModel.IsSOS), time.Now(), inputModel.UserId).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}
