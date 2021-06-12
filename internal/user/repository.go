package user

import (
	"cproject/internal/models"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Email string         `json:"email" db:"email"`
	Id    int            `json:"id" db:"uid"`
	Name  sql.NullString `json:"name"`
}

type Users []UserInfo

type UserResource struct {
	Rid       int       `json:"rid" db:"rid"`
	Uid       int       `json:"uid" db:"uid"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{db}
}

func (repo *repository) isRegisteredUser(user *models.User) (*models.User, error) {
	fmt.Println("useremail", user.Email)
	query := fmt.Sprintf("Select * from user_info where email='%s'", user.Email)

	var fetchedUsers Users
	err := repo.db.Select(&fetchedUsers, query)
	if err != nil {
		return nil, err
	}
	if len(fetchedUsers) == 0 {
		return nil, nil
	}
	fetchedUser := fetchedUsers[0]
	fmt.Println("info", fetchedUser.Email, user.Email, fetchedUser.Email == user.Email)
	if fetchedUser.Email == user.Email {
		user := &models.User{ID: fetchedUser.Id}
		return user, nil
	}
	return nil, nil
}

func (repo *repository) registerUser(user *models.User) (*models.User, error) {

	query := `INSERT INTO user_info(email) values(:email) returning uid`
	rows, err := repo.db.NamedQuery(query, user)

	if err != nil {
		return nil, err
	}

	var Id int
	if rows.Next() {
		rows.Scan(&Id)
	}
	user.ID = Id
	return user, nil
}

func (repo *repository) updateLikes(userResource *UserResource) error {

	var resourceId []int
	selectQuery := fmt.Sprintf("Select rid from user_resource where uid=%d and rid=%d", userResource.Uid, userResource.Rid)
	err := repo.db.Select(&resourceId, selectQuery)
	if err != nil {
		log.Println("error fetching from user resource", err.Error())
		return err
	}
	if len(resourceId) > 0 {
		res, err := repo.db.Exec(fmt.Sprintf("Delete from user_resource where rid=%d and uid=%d", userResource.Rid, userResource.Uid))
		if err != nil {
			log.Println("Failed to delete like", err.Error())
			return err
		}
		fmt.Println("res,", res)
	} else {
		userResource.UpdatedAt = time.Now().UTC()
		query := `INSERT INTO user_resource(uid, rid, updated_at) values(:uid, :rid,:updated_at)  `
		_, err = repo.db.NamedQuery(query, userResource)
		if err != nil {
			log.Println("error adding to user_resource", err.Error())
			return err
		}

	}
	return nil
}
