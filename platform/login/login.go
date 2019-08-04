package login

import (
	"LoginApp/platform/signup"
	"database/sql"
	"fmt"
)

func GetUserByUsername(email string, db *sql.DB) (signup.Data, error) {
	var data signup.Data
	row := db.QueryRow("select * from users where email= ?;", email)

	err := row.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password, &data.Create, &data.Update)
	if err != nil {
		fmt.Print("No details found", err.Error())
	}

	return data, err

}
