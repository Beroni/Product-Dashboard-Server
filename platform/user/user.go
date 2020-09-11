package user

type Getter interface {
	GetAll() []User
}

type Adder interface {
	Add(user User)
}

// User Nodel
type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"passwword"`
}

// Repo Nodel
type Repo struct {
	Users []User
}

// New Repo Constructor
func New() *Repo {
	return &Repo{
		Users: []User{},
	}
}

// Add User to database
func (r *Repo) Add(user User) {
	r.Users = append(r.Users, user)
}

// GetAll Users from database
func (r *Repo) GetAll() []User {
	return r.Users
}
