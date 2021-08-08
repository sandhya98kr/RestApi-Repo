package entity

import (
	//GORM is ORM library for dealing with relational databases,and it provides First,Take,Last methods
	"gorm.io/gorm"
)

//creating a struct called User
type User struct {
	//which includes fields ID,CreatedAt,UpdatedAt,DeletedAt
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	//if user is successfully created it returs nil or else it returs err
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Save(User)
	//db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
