package user

type Repository interface {
    Create(user *User) (*User, error)
    FindAll() ([]User, error)
    FindByID(id int) (*User, error)
    Update(id int, user *User) (*User, error)
    Delete(id int) error
}
