package models

import (
	"encoding/json"
	"fmt"
	mongodb "go-simple-app/db/mongo"
	"go-simple-app/utils"
	"strings"

	"github.com/fatih/structs"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

// User model
type User struct {
	Name     *string `json:"name", mapstructure:"name"`
	Username *string `json:"username" mapstructure:"username"`
	Password *string `json:"password" mapstructure:"password"`
	Age      *int    `json:"age" mapstructure:"age"`
	EmailId  *string `json:"email_id" mapstructure:"email_id"`
}

//IsValidLogin validates the user is valid for login request or not
func (user *User) IsValidLogin() bool {
	if user.Username == nil || user.Password == nil {
		return false
	}
	return true
}

//IsValidCreate validates the user is valid for create request or not
func (user *User) IsValidCreate() bool {
	if user.Username == nil || user.Password == nil || user.Age == nil || user.EmailId == nil {
		return false
	}
	return true
}

func (user *User) Create() error {
	_, err := mongodb.GetCollection().InsertOne(context.TODO(), user)
	return err

}

// GetToken generates the token using username and password
func (user *User) GetToken() string {
	return utils.GenerateToken(*user.Username + "---" + *user.Password)
}

// DecodeToken decodes the token and set username and password to user
func (user *User) DecodeToken(token string) error {
	sDec := utils.DecodeToken(token)

	splits := strings.Split(sDec, "---")
	if len(splits) != 2 {
		return fmt.Errorf("Invalid token, failed to parse")
	}

	user.Username = &splits[0]
	user.Password = &splits[1]
	return nil
}

// internal method to update the user struct
func (user *User) update(updatedUser *User) {
	if updatedUser.Name != nil {
		user.Name = updatedUser.Name
	}
	if updatedUser.Age != nil {
		user.Age = updatedUser.Age
	}
	if updatedUser.EmailId != nil {
		user.EmailId = updatedUser.EmailId
	}
	if updatedUser.Password != nil {
		user.Password = updatedUser.Password
	}
}

//UpdateProfile will update all the non-nil column from updatedUser
func (user *User) UpdateProfile(updatedUser *User) error {
	err := user.GetFullProfile()
	if err != nil {
		return err
	}
	user.update(updatedUser)
	updatedDocument := bson.M{
		"$set": user,
	}
	return mongodb.GetCollection().FindOneAndUpdate(context.TODO(), bson.D{{"username", user.Username}}, updatedDocument).Decode(user)
}

// DeleteProfile deletes the record from DB
func (user *User) DeleteProfile() error {
	_, err := mongodb.GetCollection().DeleteOne(context.TODO(), bson.D{{"username", user.Username}})
	return err
}

// GetFullProfile updates the struct with all fields from DB
func (user *User) GetFullProfile() error {
	return mongodb.GetCollection().FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(user)
}

// GetJSON returns the json object without password
func (user *User) GetJSON() ([]byte, error) {
	m := structs.Map(user)
	delete(m, "Password")
	return json.Marshal(m)
}
