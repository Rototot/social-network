package users

type UserID int64

type HashedPassword string

type User struct {
	ID        UserID
	Email     string
	Password  HashedPassword
	FirstName string
	LastName  string
	Age       int8
	Gender    Gender
	City      string
	Interests []string
}
