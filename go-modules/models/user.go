package models

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	// As we see we can declare an array of astruct this way
	users []*User
	// we dont need to initialize values here as you can see
	nextID = 1
)

// manipulate app state
func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}