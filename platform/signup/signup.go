package signup

import (
	"database/sql"
	"fmt"
)

type Getter interface {
	GetAllRegisterdInDB(db *sql.DB) []Data
}

//for in memory storage
type User struct {
	First_name string
	Last_name  string
	Email      string
	Password   string
}
type Repo struct {
	Users []User
}

//for database storage
type Data struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Create    string `json:"create"`
	Update    string `json:"update"`
}

type List struct {
	Users []Data
}

func New() *Repo {
	user1 := User{"shon", "phand", "shonphand@gmail.com", "$2a$14$z4dfE/2qNWf4aAbVnH9mY.yzj612LQrbiJpBpKjUj.UxGxFsbLMvy"}
	user2 := User{"shon", "phand", "shon@gmail.com", "$2a$14$fhEV1KyAcE1mVy84.UmFpebJSiusUI5eD/3I5sgEPtKgUa5gAtmVi"}
	return &Repo{Users: []User{user2, user1}}
}
func (r *Repo) Register(user User) {
	r.Users = append(r.Users, user)
}

func RegisterInDB(newuser Data, db *sql.DB) error {

	stmt, err := db.Prepare("insert into users (firstname,lastname,email,password ,creation_time,updation_time) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(newuser.Firstname, newuser.Lastname, newuser.Email, newuser.Password, newuser.Create, newuser.Update)

	if err != nil {
		fmt.Print(err.Error())
	}
	return err
}

func GetAllRegisterdInDB(db *sql.DB) []Data {

	var data Data
	var allusers List
	rows, err := db.Query("select * from users;")
	defer rows.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password, &data.Create, &data.Update)
		allusers.Users = append(allusers.Users, data)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	fmt.Println("data received : ", allusers.Users)
	return allusers.Users
}
