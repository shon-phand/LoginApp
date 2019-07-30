package signup

type User struct {
	First_name string
	Last_name  string
	Email      string
	Password   string
}

type Repo struct {
	Users []User
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Register(user User) {
	r.Users = append(r.Users, user)
}

func (r *Repo) GetAllRegisterd() []User {
	return r.Users
}
