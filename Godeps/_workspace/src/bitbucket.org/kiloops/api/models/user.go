package models

import (
	"time"

	"bitbucket.org/kiloops/api/ierrors"
	"bitbucket.org/kiloops/api/utils"

	"github.com/jinzhu/gorm"
	"github.com/mrkaspa/go-helpers"
)

//User model
type User struct {
	ID        int    `gorm:"primary_key"`
	Email     string `sql:"unique"`
	Active    bool   `sql:"default:0"`
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//UserLogin model
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

//UserLogged model
type UserLogged struct {
	ID     int    `gorm:"primary_key"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Active bool   `json:"active"`
}

//BeforeCreate callback
func (u *User) BeforeCreate() error {
	u.Password = helpers.MD5(u.Password)
	u.Token = helpers.RandomString(20)
	if utils.IsTest() {
		u.Active = true
	}
	return nil
}

//LoggedIn validtes if a user is logged
func (u User) LoggedIn(login UserLogin) bool {
	return helpers.MD5(login.Password) == u.Password
}

//AllProjects loads user projects
func (u User) AllProjects() *[]UsersProjects {
	userProjects := []UsersProjects{}
	Gdb.Where("user_id = ?", u.ID).Find(&userProjects)
	fullUserProjects := []UsersProjects{}
	for _, p := range userProjects {
		Gdb.Model(&p).Related(&(p.Project))
		fullUserProjects = append(fullUserProjects, p)
	}
	return &fullUserProjects
}

//OwnedProjects that the user created or is admin
func (u User) OwnedProjects() []UsersProjects {
	userProjects := []UsersProjects{}
	Gdb.Where("user_id = ? and role = ?", u.ID, Creator).Find(&userProjects)
	return userProjects
}

//CollaboratorProjects that the user can access to
func (u User) CollaboratorProjects() []UsersProjects {
	userProjects := []UsersProjects{}
	Gdb.Where("user_id = ? and role = ?", u.ID, Collaborator).Find(&userProjects)
	return userProjects
}

//CreateProject a new one
func (u User) CreateProject(txn *gorm.DB, project *Project) error {
	if txn.Create(&project).Error != nil {
		return ierrors.ErrProjectNotSaved
	}
	// Creates a relation between the user and the project
	if err := project.AddUser(txn, &u, Creator); err != nil {
		return err
	}
	return nil
}

//LeaveProject by ID
func (u User) LeaveProject(txn *gorm.DB, projectID int) error {
	var userProject UsersProjects
	txn.Where("user_id = ? and project_id = ?", u.ID, projectID).Find(&userProject)
	return txn.Delete(&userProject).Error
}

//HasAdminAccessTo by ID
func (u User) HasAdminAccessTo(projectID int) bool {
	var count int
	Gdb.Model(UsersProjects{}).Where("user_id = ? and project_id = ? and role = ?", u.ID, projectID, Creator).Count(&count)
	return count > 0
}

//HasCollaboratorAccessTo by ID
func (u User) HasCollaboratorAccessTo(projectID int) bool {
	var count int
	Gdb.Model(UsersProjects{}).Where("user_id = ? and project_id = ? and role = ?", u.ID, projectID, Collaborator).Count(&count)
	return count > 0
}

//HasWriteAccessTo By ID
func (u User) HasWriteAccessTo(projectID int) bool {
	var count int
	Gdb.Model(UsersProjects{}).Where("user_id = ? and project_id = ? and role in (?,?)", u.ID, projectID, Creator, Collaborator).Count(&count)
	return count > 0
}

//FindUser by ID
func FindUser(id int) (*User, error) {
	var user User
	Gdb.First(&user, id)
	if user.ID == 0 {
		return nil, ierrors.ErrUserNotFound
	}
	return &user, nil
}

//FindUserByEmail by email
func FindUserByEmail(email string) (*User, error) {
	var user User
	Gdb.Where("email like ?", email).First(&user)
	if user.ID == 0 {
		return nil, ierrors.ErrUserNotFound
	}
	return &user, nil
}
