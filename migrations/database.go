package migrations

import (
	"github.com/Moonieh/banking/helpers"
	"github.com/Moonieh/banking/models"
)

func createAccounts() {
	conf := GetConf()
	DB := ConnectDB(conf)

	users := [2]models.User{
		{
			Username: "Moonieh",
			Email:    "moonieh@hotmail.com",
			Password: "test12345",
		},
		{
			Username: "Loone",
			Email:    "Loone@hotmail.com",
			Password: "passpass123",
		},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Password))
		user := models.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatedPassword,
		}

		db.Create(&user)
	}
}
